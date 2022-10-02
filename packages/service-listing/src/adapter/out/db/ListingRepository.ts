import Client from "serverless-mysql";

import { ListingMapper } from "../../in/mapper";
import { ListingRepository } from "../../../application/port";
import { LoggerUtil } from "../../../utils";
import { ListingInterface, ListingRawInterface } from "../../../types";

export class ListingRepositoryImpl implements ListingRepository {
  public readonly tableName: string;
  private _client: any;
  private _logger: LoggerUtil;

  private constructor(client: any, tableName: string, className: string) {
    this._client = client;
    this.tableName = tableName;
    this._logger = new LoggerUtil(className);
  }
  public static async create(): Promise<ListingRepositoryImpl> {
    var client = Client({
      config: {
        host: process.env.DB_HOST,
        database: process.env.DB_NAME,
        user: process.env.DB_USERNAME,
        password: process.env.DB_PASSWORD,
      },
    });
    return new ListingRepositoryImpl(client, "listing", "ListingRepository");
  }

  public async setup(): Promise<void> {
    await this._client.query(
      `
        CREATE TABLE IF NOT EXISTS listing (
          id INT AUTO_INCREMENT NOT NULL, 
          host_id INT, 
          street_address VARCHAR(100) NOT NULL, 
          latitude DECIMAL(11,7) NOT NULL, 
          longitude DECIMAL(11,7) NOT NULL, 
          PRIMARY KEY (id)
        )
      `
    );
    await this._client.query(
      `
        CREATE TABLE IF NOT EXISTS image (
          id INT AUTO_INCREMENT NOT NULL,
          url VARCHAR(100),
          listing_id INT NOT NULL,
          FOREIGN KEY (listing_id) REFERENCES listing(id),
          PRIMARY KEY (id)
        )
      `
    );
  }

  public async save(data: Omit<ListingInterface, "id">): Promise<void> {
    this._logger.info(data, "save()");
    try {
      const result = await this._client.query(
        `INSERT INTO listing (host_id, street_address, latitude, longitude) VALUES(?,?,?,?)`,
        [data.hostId, data.streetAddress, data.latitude, data.longitude]
      );
      for await (const imageUrl of data.imageUrls) {
        await this._client.query(
          `INSERT INTO image (url, listing_id) VALUES (?,?)`,
          [imageUrl, result.insertId]
        );
      }
    } catch (err) {
      this._logger.error(err, "save()");
      throw err;
    }
  }

  public async addItemToListing(
    id: number,
    itemId: number
  ): Promise<{ insertId: number } | undefined> {
    this._logger.info({ id, itemId }, "addItemToListing()");
    this._logger.info(
      {
        config: {
          host: process.env.DB_HOST,
          database: process.env.DB_NAME,
          user: process.env.DB_USERNAME,
          password: process.env.DB_PASSWORD,
        },
      },
      "test()"
    );

    try {
      return await this._client.query(
        `INSERT INTO listing_item (listing_id, item_id) VALUES (?,?)`,
        id,
        itemId
      );
    } catch (err) {
      this._logger.error(err, "addItemToListing()");
      throw err;
    }
  }

  public async delete(id: number): Promise<void> {
    this._logger.info({ id }, "delete()");
    return await this._client.query(`DELETE FROM listing WHERE id = ?`, [id]);
  }

  public async findOneById(id: number): Promise<ListingInterface> {
    this._logger.info({ id }, "delete()");
    const result = await this._client.query(
      `SELECT * FROM listing WHERE id = ?`,
      [id]
    );
    return ListingMapper.toDTOFromRaw(result[0]);
  }

  public async findManyByLatLng(
    latitude: number,
    longitude: number,
    range: number = 100
  ): Promise<ListingInterface[]> {
    this._logger.info({ latitude, longitude, range }, "delete()");
    const result: ListingRawInterface[] = await this._client.query(
      `
      SELECT id, ( 3959 * acos( cos( radians(?) ) * cos( radians( latitude ) ) * cos( radians( longitude ) - radians(?) ) + sin( radians(?) ) * sin( radians( latitude ) ) ) ) AS distance, latitude, longitude FROM listing HAVING distance < ? ORDER BY distance LIMIT 0 , 20
    `,
      [latitude, longitude, latitude, range]
    );

    return result.map((item) => ListingMapper.toDTOFromRaw(item));
  }
}
