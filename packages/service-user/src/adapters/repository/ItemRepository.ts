import { ItemRepository } from "../../app/port";
import { Item } from "../../domain/model";
import { ItemRawInterface } from "../../types";
import { ItemMapper } from "../mappers";
import { AbstractRepositoryImpl } from "./AbstractRepository";

export class ItemRepositoryImpl
  extends AbstractRepositoryImpl<Item>
  implements ItemRepository
{
  public static async create(): Promise<ItemRepositoryImpl> {
    return new ItemRepositoryImpl("item", "ItemRepository");
  }

  public async setup(): Promise<void> {
    this._logger.info({}, "setup()");
    const client = this.getDBClient();
    try {
      await client.connect();
      await client.query(
        `
        CREATE TABLE IF NOT EXISTS item (
          id SERIAL NOT NULL PRIMARY KEY, 
          name VARCHAR(20) NOT NULL, 
          owner_id VARCHAR(30) NOT NULL, 
          listing_id VARCHAR(30) NOT NULL,
          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
          updated_at TIMESTAMP,
          CONSTRAINT fk_user
            FOREIGN KEY(owner_id) 
	            REFERENCES user(id)
	              ON DELETE CASCADE
        )
      `
      );

      await client.query(
        `
        CREATE TABLE IF NOT EXISTS item_images (
          id SERIAL NOT NULL PRIMARY KEY,
          url VARCHAR(150) NOT NULL,
          item_id INT NOT NULL,
          CONSTRAINT fk_item
            FOREIGN KEY(item_id) 
	            REFERENCES item(id)
	              ON DELETE CASCADE
        )
      `
      );
      await client.end();
    } catch (err) {
      await client.end();
      this._logger.error(err, "setup()");
      throw err;
    }
  }

  public async save(data: Item): Promise<Item> {
    this._logger.info(data, "save()");
    const client = this.getDBClient();

    const operations = async (data: Item | string) => {
      if (!Item.isItem(data)) {
        throw new Error("Provided data is not Item model");
      }
      const result = await client.query(
        `INSERT INTO item (name, owner_id, listing_id, created_at) VALUES($1, $2, $3, $4) RETURNING *`,
        [data.name, data.ownerId, data.listingId, data.createdAt]
      );
      for (const imageUrl of data.imageUrls) {
        await client.query(
          `INSERT INTO item_images (item_id, image_url) VALUES($1, $2)`,
          [result.rows[0].id, imageUrl]
        );
      }
    };
    const result = await this.startTransaction(operations, client, data);
    data.id = result.rows[0].id;
    return data;
    // try {
    //   await client.connect();
    //   const result = await client.query(
    //     `INSERT INTO item (name, owner_id, listing_id, created_at) VALUES($1, $2, $3, $4) RETURNING *`,
    //     [data.name, data.ownerId, data.listingId, data.createdAt]
    //   );
    //   for (const imageUrl of data.imageUrls) {
    //     await client.query(
    //       `INSERT INTO item_images (item_id, image_url) VALUES($1, $2)`,
    //       [result.rows[0].id, imageUrl]
    //     );
    //   }
    //   await client.end();
    //   return result.rows[0];
    // } catch (err) {
    //   await client.end();
    //   this._logger.error(err, "save()");
    //   throw err;
    // }
  }

  public async delete(id: string): Promise<Item> {
    this._logger.info({ id }, "delete()");
    const client = this.getDBClient();
    const operations = async () => {
      const result = await client.query(
        `DELETE FROM item WHERE id = $1 RETURNING *`,
        [id]
      );
      await client.query("DELETE FROM item_images WHERE item_id = $1", [id]);
      return result;
    };
    // try {
    //   await client.connect();
    //   const result = await client.query(
    //     `DELETE FROM item WHERE id = $1 RETURNING *`,
    //     [id]
    //   );
    //   await client.query("DELETE FROM item_images WHERE item_id = $1", [id]);
    //   await client.end();
    //   return result.rows[0];
    // } catch (err) {
    //   await client.end();
    //   this._logger.error(err, "delete()");
    //   throw err;
    // }
    const result = await this.startTransaction(operations, client, id);
    return ItemMapper.toEntityFromRaw(result.rows[0]);
  }

  public async findOneById(id: string): Promise<Item> {
    this._logger.info({ id }, "findOneById()");
    const client = this.getDBClient();
    try {
      await client.connect();
      // use left join for this
      let result = await client.query(
        `SELECT * FROM item WHERE id = $1 RETURNING *`,
        [id]
      );
      const imageUrls = await client.query(
        "SELECT * FROM item_images WHERE item_id = $1",
        [id]
      );
      const raw: ItemRawInterface = {
        ...result.rows[0],
        imageUrls,
      };
      await client.end();
      return ItemMapper.toEntityFromRaw(raw);
    } catch (err) {
      await client.end();
      this._logger.error(err, "findOneById()");
      throw err;
    }
  }
}
