import Client from "serverless-mysql";
import {
  ListingInterface,
  RDSRepository,
} from "@rental-storage-project/common";
import { ListingMapper } from "../mapper";
import { ListingRawInterface } from "../entity";

export class ListingRepository extends RDSRepository {
  public static async create(): Promise<ListingRepository> {
    var client = Client({
      config: {
        host: process.env.MYSQL_HOST,
        database: process.env.DB_NAME,
        user: process.env.USERNAME,
        password: process.env.PASSWORD,
      },
    });
    return new ListingRepository(client, "listing");
  }

  public async setup(): Promise<void> {
    await this._client.query(
      "CREATE TABLE IF NOT EXISTS listing (id INT AUTO_INCREMENT,host_id INT,street_address VARCHAR(100) NOT NULL, latitude INT NOT NULL, longitude INT NOT NULL, PRIMARY KEY (id))"
    );
  }

  public async save(data: Omit<ListingInterface, "id">): Promise<void> {
    const result = await this._client.query(
      `INSERT INTO ${this.tableName} (host_id, street_address, latitude, longitude) VALUES(?,?,?,?)`,
      [data.hostId, data.streetAddress, data.latitude, data.longitude]
    );
    return result;
  }

  public async delete(id: string): Promise<void> {
    return await this._client.query(
      `DELETE FROM ${this.tableName} WHERE id = ?`,
      [id]
    );
  }

  public async findOneById(id: string): Promise<ListingInterface> {
    const result = await this._client.query(
      `SELECT * FROM ${this.tableName} WHERE id = ?`,
      [id]
    );
    return ListingMapper.toDTOFromRaw(result[0]);
  }

  public async findManyByLatLng(
    latitude: number,
    longitude: number,
    range: number = 100
  ): Promise<ListingInterface[]> {
    const result: ListingRawInterface[] = await this._client.query(
      `
      SELECT id, ( 3959 * acos( cos( radians(?) ) * cos( radians( latitude ) ) * cos( radians( longitude ) - radians(?) ) + sin( radians(?) ) * sin( radians( latitude ) ) ) ) AS distance, latitude, longitude FROM listing HAVING distance < ? ORDER BY distance LIMIT 0 , 20
    `,
      [latitude, longitude, latitude, range]
    );
    console.log("RESULT: ", result);
    return result.map((item) => ListingMapper.toDTOFromRaw(item));
  }
}
