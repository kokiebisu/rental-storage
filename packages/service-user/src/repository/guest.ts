import Client from "serverless-mysql";
import {
  GuestInterface,
  RDSRepository,
  StorageItemInterface,
} from "@rental-storage-project/common";
import { GuestMapper } from "../mapper";

export class GuestRepository extends RDSRepository {
  public static async create(): Promise<GuestRepository> {
    var client = Client({
      config: {
        host: process.env.MYSQL_HOST,
        database: process.env.DB_NAME,
        user: process.env.DB_USERNAME,
        password: process.env.DB_PASSWORD,
      },
    });
    return new GuestRepository(client, "guest", "GuestRepository");
  }

  public async setup(): Promise<void> {
    await this._client.query(
      "CREATE TABLE IF NOT EXISTS guest (id int AUTO_INCREMENT,first_name varchar(20) NOT NULL DEFAULT '', last_name varchar(20) NOT NULL DEFAULT '', PRIMARY KEY (id))"
    );
    await this._client.query(
      "CREATE TABLE IF NOT EXISTS guest_item (guest_id INT NOT NULL, item_id INT NOT NULL, PRIMARY KEY(guest_id, item_id))"
    );
  }

  public async save(data: GuestInterface): Promise<GuestInterface> {
    const result = await this._client.query(
      `INSERT INTO guest (first_name, last_name) VALUES(?,?)`,
      [data.firstName, data.lastName]
    );
    return result;
  }

  public async delete(id: string): Promise<GuestInterface> {
    const result = await this._client.query(`DELETE FROM guest WHERE id = ?`, [
      id,
    ]);
    return result;
  }

  public async findOneById(id: string): Promise<GuestInterface> {
    const result = await this._client.query(
      `SELECT * FROM guest WHERE id = ?`,
      [id]
    );

    //     SELECT g.*, i.* FROM person p
    // INNER JOIN person_fruit pf
    // ON pf.person_id = p.id
    // INNER JOIN fruits f
    // ON f.fruit_name = pf.fruit_name
    return GuestMapper.toDTOFromRaw(result[0]);
  }

  public async findAllItemIdsByUserId(guestId: string) {
    const result = await this._client.query(
      `SELECT * FROM guest_item WHERE guest_id = ?`,
      [guestId]
    );
    return result;
  }

  public async updateStoringItem(
    userId: string,
    items: StorageItemInterface[]
  ): Promise<void> {
    for (const item of items) {
      if (!item.id) {
        throw new Error("attribute id doesn't exist");
      }
      await this._client.query(
        `INSERT INTO guest_item (guest_id, item_id) VALUES (?,?)`,
        [userId, item.id]
      );
    }
  }
}
