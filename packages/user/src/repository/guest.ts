import Client from "serverless-mysql";
import { GuestInterface } from "../entity";
import { GuestMapper } from "../mapper";
import { RDSRepository } from "./rds";

export class GuestRepository extends RDSRepository {
  public static async create(): Promise<GuestRepository> {
    var client = Client({
      config: {
        host: process.env.MYSQL_HOST,
        database: process.env.DB_NAME,
        user: process.env.USERNAME,
        password: process.env.PASSWORD,
      },
    });
    return new GuestRepository(client, "guest");
  }

  public async setup(): Promise<void> {
    await this._client.query(
      "CREATE TABLE IF NOT EXISTS host (id int AUTO_INCREMENT,first_name varchar(20) NOT NULL DEFAULT '', last_name varchar(20) NOT NULL DEFAULT '', PRIMARY KEY (id))"
    );
  }

  public async save(data: GuestInterface): Promise<GuestInterface> {
    const result = await this._client.query(
      `INSERT INTO ${this.tableName} (first_name, last_name) VALUES(?,?)`,
      [data.firstName, data.lastName]
    );
    console.log("SAVE RESULT: ", result);
    return result;
  }

  public async delete(id: string): Promise<GuestInterface> {
    const result = await this._client.query(
      `DELETE FROM ${this.tableName} WHERE id = ?`,
      [id]
    );
    console.log("DELETE RESULT: ", result);
    return result;
  }

  public async findOneById(id: string): Promise<GuestInterface> {
    const result = await this._client.query(
      `SELECT * FROM ${this.tableName} WHERE id = ?`,
      [id]
    );
    console.log("GETBYID RESULT: ", result);
    return GuestMapper.toDTOFromRaw(result[0]);
  }
}
