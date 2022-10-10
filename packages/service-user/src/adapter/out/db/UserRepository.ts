import Client from "serverless-mysql";

import { LoggerUtil } from "../../../utils";
import { UserRepository } from "../../../application/port";
import { UserMapper } from "../../in/mapper";
import { User } from "../../../domain/model";

export class UserRepositoryImpl implements UserRepository {
  public readonly tableName: string;
  private _client: any;
  private _logger: LoggerUtil;

  private constructor(client: any, tableName: string, className: string) {
    this._client = client;
    this.tableName = tableName;
    this._logger = new LoggerUtil(className);
  }
  public static async create(): Promise<UserRepositoryImpl> {
    var client = Client({
      config: {
        host: process.env.DB_HOST,
        database: process.env.DB_NAME,
        user: process.env.DB_USERNAME,
        password: process.env.DB_PASSWORD,
      },
    });
    return new UserRepositoryImpl(client, "user", "UserRepository");
  }

  public async setup(): Promise<void> {
    await this._client.query(
      `
        CREATE TABLE IF NOT EXISTS user (
          id int AUTO_INCREMENT, 
          uid VARCHAR(32), 
          first_name varchar(20) NOT NULL DEFAULT '', 
          last_name varchar(20) NOT NULL DEFAULT '', 
          email_address VARCHAR(50) NOT NULL, 
          password VARCHAR(50) NOT NULL,
          created_at DATE NOT NULL, 
          updated_at DATE, 
          UNIQUE (email_address), 
          PRIMARY KEY (id)
        )
      `
    );
  }

  public async save(data: User): Promise<User> {
    this._logger.info(data, "save()");
    try {
      const result = await this._client.query(
        `INSERT INTO user (uid, email_address, password, first_name, last_name, created_at) VALUES(?,?,?,?,?,?)`,
        [
          data.uid,
          data.emailAddress.value,
          data.password,
          data.firstName,
          data.lastName,
          data.createdAt,
        ]
      );

      data.id = result.insertId;

      return data;
    } catch (err) {
      this._logger.error(err, "save()");
      throw err;
    }
  }

  public async delete(id: number): Promise<void> {
    this._logger.info(id, "delete()");
    try {
      // set up commit/transaction
      const result = await this._client.query(`DELETE FROM user WHERE id = ?`, [
        id,
      ]);
      console.log("DELETE RESULT: ", result);
      // return result
    } catch (err) {
      this._logger.error(err, "delete()");
    }
  }

  public async findOneById(id: number): Promise<User> {
    this._logger.info({ id }, "findOneById()");

    const result = await this._client.query(
      `
        SELECT user.*, payment.id AS payment_id, payment.provider_id AS payment_provider_id, payment.provider_type AS payment_provider_type FROM user 
        INNER JOIN payment ON user.id = payment.user_id 
        WHERE user.id = ?
      `,
      [id]
    );

    return UserMapper.toEntityFromRaw(result[0]);
  }

  public async findOneByEmail(emailAddress: string): Promise<User> {
    this._logger.info({ emailAddress }, "findOneByEmail()");
    try {
      const result = await this._client.query(
        `
          SELECT user.*, payment.id AS payment_id, payment.provider_id AS payment_provider_id, payment.provider_type AS payment_provider_type FROM user 
          INNER JOIN payment ON user.id = payment.user_id 
          WHERE user.email_address = ?`,
        [emailAddress]
      );

      console.log("FindOneByEmail: ", result);

      return UserMapper.toEntityFromRaw(result[0]);
    } catch (err) {
      this._logger.error(err, "findOneByEmail()");
      throw err;
    }
  }
}
