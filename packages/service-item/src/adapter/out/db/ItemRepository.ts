import Client from "serverless-mysql";
import {
  LoggerService,
  StorageItemInterface,
} from "@rental-storage-project/common";
import { StorageItemRepository } from "../../../application/port";
import { StorageItemMapper } from "../../in/mapper";

export class StorageItemRepositoryImpl implements StorageItemRepository {
  public readonly tableName: string;
  private _client: any;
  private _logger: LoggerService;

  private constructor(client: any, tableName: string, className: string) {
    this._client = client;
    this.tableName = tableName;
    this._logger = new LoggerService(className);
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
      "CREATE TABLE IF NOT EXISTS storage (id int AUTO_INCREMENT,name varchar(20) NOT NULL, user_id varchar(30) NOT NULL, listing_id varchar(30) NOT NULL, created_at varchar(30), updated_at varchar(30), PRIMARY KEY (id))"
    );
  }

  public async save(data: StorageItemInterface): Promise<{ insertId: number }> {
    this._logger.info(data, "save()");
    return await this._client.query(
      `INSERT INTO storage (name, user_id, listing_id, created_at) VALUES(?,?,?,?)`,
      [data.name, data.userId, data.listingId, data.createdAt]
    );
  }

  public async delete(id: string): Promise<StorageItemInterface> {
    const result = await this._client.query(
      `DELETE FROM storage WHERE id = ?`,
      [id]
    );
    return result;
  }

  public async findOneById(id: string): Promise<StorageItemInterface> {
    const result = await this._client.query(
      `SELECT * FROM storage WHERE id = ?`,
      [id]
    );
    return StorageItemMapper.toDTOFromRaw(result[0]);
  }
}
