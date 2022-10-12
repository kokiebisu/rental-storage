import { Client } from "pg";

import { LoggerUtil } from "../../../utils";
import { PaymentRepository } from "../../../application/port";
import { Payment } from "../../../domain/model";

export class PaymentRepositoryImpl implements PaymentRepository {
  public readonly tableName: string;
  private _logger: LoggerUtil;

  private constructor(tableName: string, className: string) {
    this.tableName = tableName;
    this._logger = new LoggerUtil(className);
  }
  public static async create(): Promise<PaymentRepositoryImpl> {
    return new PaymentRepositoryImpl("payment", "PaymentRepository");
  }

  public getDBClient() {
    const client = new Client({
      host: process.env.DB_HOST,
      database: process.env.DB_NAME,
      user: process.env.DB_USERNAME,
      password: process.env.DB_PASSWORD,
      port: 5432,
    });
    return client;
  }

  public async setup(): Promise<void> {
    this._logger.info({}, "setup()");
    const client = this.getDBClient();
    try {
      await client.connect();
      await client.query(
        `
          CREATE TABLE IF NOT EXISTS payment (
            id SERIAL NOT NULL PRIMARY KEY,
            user_id INT NOT NULL,
            provider_id VARCHAR(32) UNIQUE,
            provider_type VARCHAR(10),
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
            updated_at TIMESTAMP,
            CONSTRAINT fk_user
              FOREIGN KEY(user_id) 
                REFERENCES user_account(id)
                  ON DELETE CASCADE
          )
        `
      );
      await client.end();
    } catch (err) {
      await client.end();
      console.error(err);
      throw err;
    }
  }

  public async save(data: Payment): Promise<Payment> {
    this._logger.info(data, "save()");
    const client = this.getDBClient();
    try {
      await client.connect();
      const result = await client.query(
        `
          INSERT INTO payment (
          provider_id, user_id, provider_type) VALUES($1, $2, $3)
          RETURNING *
        `,
        [data.providerId, data.userId, data.providerType]
      );
      console.log("SAVED PAYMENT SUCCESSFULLY RESULT: ", result);
      data.id = result.rows[0].id;
      await client.end();
      return data;
    } catch (err) {
      this._logger.error(err, "savePayment()");
      await client.end();
      throw err;
    }
  }
}
