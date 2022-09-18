import Client from "serverless-mysql";
import { GuestInterface } from "../entity";
import { RDSRepository } from "./rds";

export class GuestRepository extends RDSRepository {
  public static async create(): Promise<GuestRepository> {
    console.log("HOST: ", process.env.MYSQL_HOST);
    console.log("Database: ", process.env.DB_NAME);
    console.log("User: ", process.env.USERNAME);
    console.log("Password: ", process.env.PASSWORD);
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

  public async save(data: GuestInterface): Promise<GuestInterface> {
    const result = await this._client.query(
      `INSERT INTO ${this.tableName} (first_name, last_name) VALUES(?,?)`,
      [data.firstName, data.lastName]
    );
    console.log("SAVE RESULT: ", result);
    return result;
  }
}
