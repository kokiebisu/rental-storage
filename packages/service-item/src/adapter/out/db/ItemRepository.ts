import Client from "serverless-mysql";

import { StorageItemRepository } from "../../../application/port";
import { StorageItemInterface } from "../../../types";
import { LoggerUtil } from "../../../utils";
import { StorageItemMapper } from "../../in/mapper";

export class StorageItemRepositoryImpl implements StorageItemRepository {
  public readonly tableName: string;
  private _client: any;
  private _logger: LoggerUtil;

  private constructor(client: any, tableName: string, className: string) {
    this._client = client;
    this.tableName = tableName;
    this._logger = new LoggerUtil(className);
  }

  public static async create(): Promise<StorageItemRepositoryImpl> {
    var client = Client({
      config: {
        host: process.env.DB_HOST,
        database: process.env.DB_NAME,
        user: process.env.DB_USERNAME,
        password: process.env.DB_PASSWORD,
      },
    });
    return new StorageItemRepositoryImpl(
      client,
      "storage",
      "StorageItemRepository"
    );
  }

  public async setup(): Promise<void> {
    await this._client.query(
      `
        CREATE TABLE IF NOT EXISTS storage (
          id int AUTO_INCREMENT, 
          name VARCHAR(20) NOT NULL, 
          guest_id VARCHAR(30) NOT NULL, 
          listing_id VARCHAR(30) NOT NULL,
          created_at DATE, 
          updated_at DATE, 
          PRIMARY KEY (id),
        )
      `
    );
    await this._client.query(
      `
        CREATE TABLE IF NOT EXISTS image (
          id int AUTO_INCREMENT,
          url VARCHAR(150),
          item_id INT NOT NULL,
          FOREIGN KEY (item_id) REFERENCES storage(item_id),
          FOREIGN KEY(person_id) REFERENCES person(id)
        )
      `
    );
  }

  public async save(
    data: StorageItemInterface
  ): Promise<{ insertId: number } | null> {
    this._logger.info(data, "save()");
    try {
      const result = await this._client.query(
        `INSERT INTO storage (name, guest_id, listing_id, created_at) VALUES(?,?,?,?)`,
        [data.name, data.guestId, data.listingId, data.createdAt]
      );
      for (const imageUrl of data.imageUrls) {
        await this._client.query(
          `INSERT INTO item_image_url (item_id, image_url) VALUES(?,?)`,
          [result.insertId, imageUrl]
        );
      }
      return result;
    } catch (err) {
      this._logger.error(err, "save()");
      return null;
    }
  }

  public async delete(id: string): Promise<StorageItemInterface> {
    const result = await this._client.query(
      `DELETE FROM storage WHERE id = ?`,
      [id]
    );
    await this._client.query("DELETE FROM item_image_url WHERE item_id = ?", [
      id,
    ]);
    return result;
  }

  public async findOneById(id: string): Promise<StorageItemInterface> {
    let result = await this._client.query(
      `SELECT * FROM storage WHERE id = ?`,
      [id]
    );
    const imageUrls = await this._client.query(
      "SELECT * FROM item_image_url WHERE item_id = ?",
      [id]
    );
    result = {
      ...result[0],
      imageUrls,
    };
    return StorageItemMapper.toDTOFromRaw(result);
  }
}
