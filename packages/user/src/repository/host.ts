import Client from "serverless-mysql";
import { HostInterface } from "../entity";
import { RDSRepository } from "./rds";

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
    return new HostRepository(client, "Host");
  }

  public async save(data: HostInterface): Promise<HostInterface> {
    const result = await this._client.query(
      "INSERT INTO users (firstName, lastName) VALUES(?,?)",
      [data.firstName, data.lastName]
    );
    return result;
  }
}
