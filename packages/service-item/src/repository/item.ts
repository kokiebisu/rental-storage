import Client from "serverless-mysql";
import {
  StorageItemInterface,
  RDSRepository,
} from "@rental-storage-project/common";
import { StorageItemMapper } from "../mapper";

export class StorageItemRepository extends RDSRepository {
  public static async create(): Promise<StorageItemRepository> {
    var client = Client({
      config: {
        host: process.env.MYSQL_HOST,
        database: process.env.DB_NAME,
        user: process.env.DB_USERNAME,
        password: process.env.DB_PASSWORD,
      },
    });
    return new StorageItemRepository(
      client,
      "storage",
      "StorageItemRepository"
    );
  }

  public async setup(): Promise<void> {
    await this._client.query(
      "CREATE TABLE IF NOT EXISTS storage (id int AUTO_INCREMENT,name varchar(20) NOT NULL, user_id varchar(30) NOT NULL, listing_id varchar(30) NOT NULL, created_at varchar(30), updated_at varchar(30), PRIMARY KEY (id))"
    );
  }

  public async save(data: StorageItemInterface): Promise<StorageItemInterface> {
    const result = await this._client.query(
      `INSERT INTO ${this.tableName} (name, user_id, listing_id, created_at) VALUES(?,?,?,?)`,
      [data.name, data.userId, data.listingId, data.createdAt]
    );
    return result;
  }

  public async delete(id: string): Promise<StorageItemInterface> {
    const result = await this._client.query(
      `DELETE FROM ${this.tableName} WHERE id = ?`,
      [id]
    );
    return result;
  }

  public async findOneById(id: string): Promise<StorageItemInterface> {
    const result = await this._client.query(
      `SELECT * FROM ${this.tableName} WHERE id = ?`,
      [id]
    );
    return StorageItemMapper.toDTOFromRaw(result[0]);
  }
}
