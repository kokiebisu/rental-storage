import Client from "serverless-mysql";

import { GuestMapper } from "../../in/mapper";
import { GuestRepository } from "../../../application/port";
import { LoggerUtil } from "../../../utils";
import { GuestInterface, StorageItemInterface } from "../../../types";

export class GuestRepositoryImpl implements GuestRepository {
  public readonly tableName: string;
  private _client: any;
  private _logger: LoggerUtil;

  private constructor(client: any, tableName: string, className: string) {
    this._client = client;
    this.tableName = tableName;
    this._logger = new LoggerUtil(className);
  }
  public static async create(): Promise<GuestRepositoryImpl> {
    var client = Client({
      config: {
        host: process.env.DB_HOST,
        database: process.env.DB_NAME,
        user: process.env.DB_USERNAME,
        password: process.env.DB_PASSWORD,
      },
    });
    return new GuestRepositoryImpl(client, "guest", "GuestRepository");
  }

  public async setup(): Promise<void> {
    await this._client.query(
      `
        CREATE TABLE IF NOT EXISTS guest (
          id int AUTO_INCREMENT, 
          uid VARCHAR(32), 
          first_name varchar(20) NOT NULL DEFAULT '', 
          last_name varchar(20) NOT NULL DEFAULT '', 
          email_address VARCHAR(50) NOT NULL, 
          password VARCHAR(50) NOT NULL,
          created_at DATE NOT NULL, 
          updated_at DATE, 
          UNIQUE (email_address), 
          PRIMARY KEY (id)
        )
      `
    );
  }

  public async save(
    data: GuestInterface
  ): Promise<{ insertId: number } | undefined> {
    this._logger.info(data, "save()");
    try {
      const result = await this._client.query(
        `INSERT INTO guest (uid, email_address, password, first_name, last_name, created_at) VALUES(?,?,?,?,?,?)`,
        [
          data.uid,
          data.emailAddress,
          data.password,
          data.firstName,
          data.lastName,
          data.createdAt,
        ]
      );
      return result;
    } catch (err) {
      this._logger.error(err, "save()");
    }
  }

  public async delete(id: number): Promise<GuestInterface> {
    this._logger.info(id, "delete()");
    const result = await this._client.query(`DELETE FROM guest WHERE id = ?`, [
      id,
    ]);
    return result;
  }

  public async findOneById(id: number): Promise<GuestInterface> {
    this._logger.info({ id }, "findOneById()");
    const result = await this._client.query(
      `SELECT * FROM guest WHERE id = ?`,
      [id]
    );

    return GuestMapper.toDTOFromRaw(result[0]);
  }

  public async findAllItemIdsByGuestId(id: number): Promise<any> {
    this._logger.info(id, "findAllItemIdsByUserId()");
    const result = await this._client.query(
      `SELECT * FROM guest_item WHERE guest_id = ?`,
      [id]
    );
    return result;
  }

  public async updateStoringItem(
    id: number,
    items: StorageItemInterface[]
  ): Promise<void> {
    this._logger.info({ id, items }, "updateStoringItem()");
    for (const item of items) {
      if (!item.id) {
        throw new Error("attribute id doesn't exist");
      }
      await this._client.query(
        `INSERT INTO guest_item (guest_id, item_id) VALUES (?,?)`,
        [id, item.id]
      );
    }
  }
}
