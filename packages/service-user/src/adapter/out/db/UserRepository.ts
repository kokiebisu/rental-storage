import Client from "serverless-mysql";

import { LoggerUtil } from "../../../utils";
import { UserInterface } from "../../../types";
import { UserRepository } from "../../../application/port";
import { UserMapper } from "../../in/mapper";
import { PaymentInterface } from "../../../types/Payment";

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

    await this._client.query(
      `
        CREATE TABLE IF NOT EXISTS payment_user (
          id int AUTO_INCREMENT,
          uid VARCHAR(32),
          user_id VARCHAR(32),
          payment_id VARCHAR(32),
          provider_type VARCHAR(10),
          created_at DATE NOT NULL, 
          updated_at DATE, 
          UNIQUE (payment_id),
          PRIMARY KEY (id)
        )
      `
    );
  }

  public async save(
    data: UserInterface
  ): Promise<{ insertId: number; uid: string }> {
    this._logger.info(data, "save()");
    try {
      const result = await this._client.query(
        `INSERT INTO user (uid, email_address, password, first_name, last_name, created_at) VALUES(?,?,?,?,?,?)`,
        [
          data.uid,
          data.emailAddress,
          data.password,
          data.firstName,
          data.lastName,
          data.createdAt,
        ]
      );

      return {
        insertId: result.insertId,
        uid: data.uid,
      };
    } catch (err) {
      this._logger.error(err, "save()");
      throw err;
    }
  }

  public async savePayment(data: PaymentInterface) {
    await this._client.query(
      `INSERT INTO payment_user (uid, payment_id, user_id, provider_type) VALUES(?,?,?,?)`,
      [data.uid, data.customerId, data.userId, data.providerType]
    );
  }

  public async delete(id: number): Promise<UserInterface> {
    this._logger.info(id, "delete()");
    const result = await this._client.query(`DELETE FROM user WHERE id = ?`, [
      id,
    ]);
    return result;
  }

  public async findOneById(id: number): Promise<UserInterface> {
    this._logger.info({ id }, "findOneById()");
    const result = await this._client.query(`SELECT * FROM user WHERE id = ?`, [
      id,
    ]);

    return UserMapper.toDTOFromRaw(result[0]);
  }

  public async findOneByEmail(
    emailAddress: string
  ): Promise<UserInterface | null> {
    this._logger.info({ emailAddress }, "findOneByEmail()");
    try {
      const result = await this._client.query(
        `SELECT * FROM user where email_address = ?`,
        [emailAddress]
      );

      return UserMapper.toDTOFromRaw(result[0]);
    } catch (err) {
      this._logger.error(err, "findOneByEmail()");
      throw err;
    }
  }
}
