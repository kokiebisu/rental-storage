import { UserRepository } from "../../app/port";
import { User } from "../../domain/model";
import { UserMapper } from "../mappers";
import { AbstractRepositoryImpl } from "./AbstractRepository";

export class UserRepositoryImpl
  extends AbstractRepositoryImpl<User>
  implements UserRepository
{
  public static async create(): Promise<UserRepositoryImpl> {
    try {
      return new UserRepositoryImpl("user", "UserRepository");
    } catch (err) {
      console.error(err);
      throw err;
    }
  }

  public async setup(): Promise<void> {
    const client = this.getDBClient();
    try {
      await client.connect();
      await client.query(
        `
          CREATE TABLE IF NOT EXISTS user_account (
            id SERIAL NOT NULL PRIMARY KEY, 
            uid VARCHAR(64) NOT NULL, 
            first_name VARCHAR(32) NOT NULL, 
            last_name VARCHAR(32) NOT NULL, 
            email_address VARCHAR(64) NOT NULL UNIQUE, 
            password VARCHAR(64) NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
            updated_at TIMESTAMP
          )
        `
      );
      await client.end();
    } catch (err) {
      try {
        await client.end();
      } catch (err) {
        console.error(err);
      }
      console.error(err);
      throw err;
    }
  }

  public async save(data: User): Promise<User> {
    this._logger.info(data, "save()");
    const client = this.getDBClient();
    const operations = async (data: User | string) => {
      if (!User.isUser(data)) {
        throw new Error("Provided data is not User model");
      }
      const result = await client.query(
        `
          INSERT INTO user_account (uid, email_address, password, first_name, last_name, created_at) 
          VALUES ($1, $2, $3, $4, $5, $6)
          RETURNING *;
        `,
        [
          data.uid,
          data.emailAddress.value,
          data.password,
          data.firstName.value,
          data.lastName.value,
          data.createdAt,
        ]
      );
      return result;
    };
    const result = await this.startTransaction(operations, client, data);
    data.id = result.rows[0].id;
    return data;
  }

  public async delete(uid: string): Promise<User> {
    this._logger.info(uid, "delete()");
    const client = this.getDBClient();
    const operations = async () => {
      const result = await client.query(
        `DELETE FROM user_account WHERE uid = $1 RETURNING *`,
        [uid]
      );
      return result;
    };
    const result = await this.startTransaction(operations, client, uid);
    return UserMapper.toEntityFromRaw(result.rows[0]);
  }

  public async findOneById(uid: string): Promise<User> {
    this._logger.info({ uid }, "findOneById()");
    const client = this.getDBClient();

    try {
      await client.connect();
      const result = await client.query(
        `
          SELECT user_account.*, payment.id AS payment_id, payment.provider_id AS payment_provider_id, payment.provider_type AS payment_provider_type FROM user_account 
          LEFT JOIN payment ON user_account.id = payment.user_id 
          WHERE user_account.uid = $1
        `,
        [uid]
      );
      await client.end();
      return UserMapper.toEntityFromRaw(result.rows[0]);
    } catch (err) {
      await client.end();
      this._logger.error(err, "findOneById()");
      throw err;
    }
  }

  public async findOneByEmail(emailAddress: string): Promise<User> {
    this._logger.info({ emailAddress }, "findOneByEmail()");
    const client = this.getDBClient();

    try {
      await client.connect();
      const result = await client.query(
        `
          SELECT user_account.*, payment.id AS payment_id, payment.provider_id AS payment_provider_id, payment.provider_type AS payment_provider_type FROM user_account 
          LEFT JOIN payment ON user_account.id = payment.user_id 
          WHERE user_account.email_address = $1
        `,
        [emailAddress]
      );

      await client.end();
      return UserMapper.toEntityFromRaw(result.rows[0]);
    } catch (err) {
      await client.end();
      this._logger.error(err, "findOneByEmail()");
      throw err;
    }
  }
}
