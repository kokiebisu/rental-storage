import Client from "serverless-mysql";

import { LoggerUtil } from "../../../utils";
import { PaymentRepository } from "../../../application/port";
import { Payment } from "../../../domain/model";

export class PaymentRepositoryImpl implements PaymentRepository {
  public readonly tableName: string;
  private _client: any;
  private _logger: LoggerUtil;

  private constructor(client: any, tableName: string, className: string) {
    this._client = client;
    this.tableName = tableName;
    this._logger = new LoggerUtil(className);
  }
  public static async create(): Promise<PaymentRepositoryImpl> {
    var client = Client({
      config: {
        host: process.env.DB_HOST,
        database: process.env.DB_NAME,
        user: process.env.DB_USERNAME,
        password: process.env.DB_PASSWORD,
      },
    });
    return new PaymentRepositoryImpl(client, "payment", "PaymentRepository");
  }

  public async setup(): Promise<void> {
    await this._client.query(
      `
        CREATE TABLE IF NOT EXISTS payment (
          id int AUTO_INCREMENT,
          user_id VARCHAR(32),
          provider_id VARCHAR(32),
          provider_type VARCHAR(10),
          created_at DATE NOT NULL, 
          updated_at DATE, 
          UNIQUE (provider_id),
          PRIMARY KEY (id)
        )
      `
    );
  }

  public async save(data: Payment): Promise<Payment> {
    this._logger.info(data, "savePayment()");
    try {
      const result = await this._client.query(
        `INSERT INTO payment (provider_id, user_id, provider_type) VALUES(?,?,?)`,
        [data.providerId, data.userId, data.providerType]
      );
      data.id = result.insertId;
      return data;
    } catch (err) {
      this._logger.error(err, "savePayment()");
      throw err;
    }
  }
}
