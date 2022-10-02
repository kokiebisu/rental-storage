import Client from "serverless-mysql";

import { HostMapper } from "../../in/mapper";
import { HostRepository } from "../../../application/port";
import { LoggerUtil } from "../../../utils";
import { HostInterface } from "../../../types";

export class HostRepositoryImpl implements HostRepository {
  public readonly tableName: string;
  private _client: any;
  private _logger: LoggerUtil;

  private constructor(client: any, tableName: string, className: string) {
    this._client = client;
    this.tableName = tableName;
    this._logger = new LoggerUtil(className);
  }
  public static async create(): Promise<HostRepositoryImpl> {
    var client = Client({
      config: {
        host: process.env.DB_HOST,
        database: process.env.DB_NAME,
        user: process.env.DB_USERNAME,
        password: process.env.DB_PASSWORD,
      },
    });
    return new HostRepositoryImpl(client, "host", "HostRepository");
  }

  public async setup(): Promise<void> {
    this._logger.info({}, "setup()");
    await this._client.query(
      `
        CREATE TABLE IF NOT EXISTS host (
          id INT AUTO_INCREMENT, 
          uid VARCHAR (32), 
          first_name VARCHAR(20) NOT NULL DEFAULT '', 
          last_name VARCHAR(20) NOT NULL DEFAULT '', 
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

  public async save(data: HostInterface): Promise<{ insertId: string }> {
    this._logger.info({ data }, "save()");
    const result = await this._client.query(
      `INSERT INTO host (uid, email_address, password,first_name, last_name, created_at) VALUES(?,?,?,?,?,?)`,
      [
        data.uid,
        data.emailAddress,
        data.password,
        data.firstName,
        data.lastName,
        data.createdAt,
      ]
    );
    return result;
  }

  public async delete(id: number): Promise<HostInterface> {
    this._logger.info({ id }, "delete()");
    const result = await this._client.query(`DELETE FROM host WHERE id = ?`, [
      id,
    ]);
    return result;
  }

  public async findOneById(id: number): Promise<HostInterface> {
    this._logger.info({ id }, "findOneById()");
    const result = await this._client.query(`SELECT * FROM host WHERE id = ?`, [
      id,
    ]);
    return HostMapper.toDTOFromRaw(result[0]);
  }
}
