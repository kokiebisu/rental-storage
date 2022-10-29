import { PaymentRepository } from "../../port";
import { Payment } from "../../domain/model";
import { AbstractRepositoryImpl } from "./AbstractRepository";

export class PaymentRepositoryImpl
  extends AbstractRepositoryImpl<Payment>
  implements PaymentRepository
{
  public static async create(): Promise<PaymentRepositoryImpl> {
    return new PaymentRepositoryImpl("payment", "PaymentRepository");
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
    const operations = async (data: Payment | string) => {
      if (!Payment.isPayment(data)) {
        throw new Error("Provided data is not Payment model");
      }
      const result = await client.query(
        `
          INSERT INTO payment (
          provider_id, user_id, provider_type) VALUES($1, $2, $3)
          RETURNING *
        `,
        [data.providerId, data.userId, data.providerType]
      );
      return result;
    };
    const result = await this.startTransaction(operations, client, data);
    data.id = result.rows[0].id;
    return result;
  }
}
