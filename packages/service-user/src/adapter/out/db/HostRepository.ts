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
    await this._client.query(
      "CREATE TABLE IF NOT EXISTS host (id int AUTO_INCREMENT,first_name varchar(20) NOT NULL DEFAULT '', last_name varchar(20) NOT NULL DEFAULT '', created_at DATETIME NOT NULL, updated_at DATETIME, PRIMARY KEY (id))"
    );
  }

  public async save(data: HostInterface): Promise<{ insertId: string }> {
    const result = await this._client.query(
      `INSERT INTO host (first_name, last_name, created_at) VALUES(?,?,?)`,
      [data.firstName, data.lastName, data.createdAt]
    );
    return result;
  }

  public async delete(id: string): Promise<HostInterface> {
    const result = await this._client.query(`DELETE FROM host WHERE id = ?`, [
      id,
    ]);
    return result;
  }

  public async findOneById(id: string): Promise<HostInterface> {
    const result = await this._client.query(`SELECT * FROM host WHERE id = ?`, [
      id,
    ]);
    return HostMapper.toDTOFromRaw(result[0]);
  }
}
