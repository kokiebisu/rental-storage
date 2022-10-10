import Client from "serverless-mysql";

import { LoggerUtil } from "../../../utils";
import { ItemRepository } from "../../../application/port";
import { ItemMapper } from "../../in/mapper";
import { Item } from "../../../domain/model";

export class ItemRepositoryImpl implements ItemRepository {
  public readonly tableName: string;
  private _client: any;
  private _logger: LoggerUtil;

  private constructor(client: any, tableName: string, className: string) {
    this._client = client;
    this.tableName = tableName;
    this._logger = new LoggerUtil(className);
  }
  public static async create(): Promise<ItemRepositoryImpl> {
    var client = Client({
      config: {
        host: process.env.DB_HOST,
        database: process.env.DB_NAME,
        user: process.env.DB_USERNAME,
        password: process.env.DB_PASSWORD,
      },
    });
    return new ItemRepositoryImpl(client, "item", "ItemRepository");
  }

  public async setup(): Promise<void> {
    await this._client.query(
      `
        CREATE TABLE IF NOT EXISTS item (
          id int AUTO_INCREMENT, 
          name VARCHAR(20) NOT NULL, 
          owner_id VARCHAR(30) NOT NULL, 
          listing_id VARCHAR(30) NOT NULL,
          created_at DATE NOT NULL, 
          updated_at DATE, 
          PRIMARY KEY (id),
        )
      `
    );

    await this._client.query(
      `
        CREATE TABLE IF NOT EXISTS item_images (
          id int AUTO_INCREMENT,
          url VARCHAR(150) NOT NULL,
          item_id INT NOT NULL,
          PRIMARY KEY (id)
        )
      `
    );
  }

  public async save(data: Item): Promise<Item> {
    this._logger.info(data, "save()");
    try {
      const result = await this._client.query(
        `INSERT INTO item (name, owner_id, listing_id, created_at) VALUES(?,?,?,?)`,
        [data.name, data.ownerId, data.listingId, data.createdAt]
      );
      for (const imageUrl of data.imageUrls) {
        await this._client.query(
          `INSERT INTO item_images (item_id, image_url) VALUES(?,?)`,
          [result.insertId, imageUrl]
        );
      }
      return result;
    } catch (err) {
      this._logger.error(err, "save()");
      throw err;
    }
  }

  public async delete(id: string): Promise<void> {
    const result = await this._client.query(`DELETE FROM item WHERE id = ?`, [
      id,
    ]);
    await this._client.query("DELETE FROM item_images WHERE item_id = ?", [id]);
    return result;
  }

  public async findOneById(id: string): Promise<Item> {
    let result = await this._client.query(`SELECT * FROM item WHERE id = ?`, [
      id,
    ]);
    const imageUrls = await this._client.query(
      "SELECT * FROM item_images WHERE item_id = ?",
      [id]
    );
    result = {
      ...result[0],
      imageUrls,
    };
    return ItemMapper.toEntityFromRaw(result);
  }
}
