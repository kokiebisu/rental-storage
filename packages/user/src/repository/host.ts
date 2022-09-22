import Client from "serverless-mysql";
import { HostInterface, RDSRepository } from "@rental-storage-project/common";
import { HostMapper } from "../mapper";

export class HostRepository extends RDSRepository {
  public static async create(): Promise<HostRepository> {
    var client = Client({
      config: {
        host: process.env.MYSQL_HOST,
        database: process.env.DB_NAME,
        user: process.env.USERNAME,
        password: process.env.PASSWORD,
      },
    });
    return new HostRepository(client, "host");
  }

  public async setup(): Promise<void> {
    await this._client.query(
      "CREATE TABLE IF NOT EXISTS host (id int AUTO_INCREMENT,first_name varchar(20) NOT NULL DEFAULT '', last_name varchar(20) NOT NULL DEFAULT '', PRIMARY KEY (id))"
    );
  }

  public async save(data: HostInterface): Promise<HostInterface> {
    const result = await this._client.query(
      `INSERT INTO ${this.tableName} (first_name, last_name) VALUES(?,?)`,
      [data.firstName, data.lastName]
    );
    return result;
  }

  public async delete(id: string): Promise<HostInterface> {
    const result = await this._client.query(
      `DELETE FROM ${this.tableName} WHERE id = ?`,
      [id]
    );
    return result;
  }

  public async findOneById(id: string): Promise<HostInterface> {
    const result = await this._client.query(
      `SELECT * FROM ${this.tableName} WHERE id = ?`,
      [id]
    );
    return HostMapper.toDTOFromRaw(result[0]);
  }
}
