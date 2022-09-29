import Client from "serverless-mysql";
import { HostInterface, LoggerService } from "@rental-storage-project/common";

import { HostMapper } from "../../in/mapper";
import { HostRepository } from "../../../application/port";

export class HostRepositoryImpl implements HostRepository {
  public readonly tableName: string;
  private _client: any;
  private _logger: LoggerService;

  private constructor(client: any, tableName: string, className: string) {
    this._client = client;
    this.tableName = tableName;
    this._logger = new LoggerService(className);
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
      "CREATE TABLE IF NOT EXISTS host (id int AUTO_INCREMENT,first_name varchar(20) NOT NULL DEFAULT '', last_name varchar(20) NOT NULL DEFAULT '', PRIMARY KEY (id))"
    );
  }

  public async save(data: HostInterface): Promise<{ insertId: string }> {
    const result = await this._client.query(
      `INSERT INTO host (first_name, last_name) VALUES(?,?)`,
      [data.firstName, data.lastName]
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
