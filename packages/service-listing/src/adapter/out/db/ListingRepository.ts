import Client from "serverless-mysql";

import { ListingMapper } from "../../in/mapper";
import { ListingRepository } from "../../../application/port";
import { LoggerUtil } from "../../../utils";
import { ListingRawInterface } from "../../../types";
import { Listing } from "../../../domain/model";

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
          lender_id VARCHAR(32), 
          street_address VARCHAR(100) NOT NULL, 
          latitude DECIMAL(11,7) NOT NULL, 
          longitude DECIMAL(11,7) NOT NULL, 
          PRIMARY KEY (id)
        )
      `
    );
    await this._client.query(
      `
        CREATE TABLE IF NOT EXISTS images_listing (
          id INT AUTO_INCREMENT NOT NULL,
          url VARCHAR(100),
          listing_id INT NOT NULL,
          FOREIGN KEY (listing_id) REFERENCES listing(id),
          PRIMARY KEY (id)
        )
      `
    );
  }

  public async save(data: Listing): Promise<Listing> {
    this._logger.info(data, "save()");
    try {
      // commit/transaction add rollback behavior
      const result = await this._client.query(
        `INSERT INTO listing (lender_id, street_address, latitude, longitude) VALUES(?,?,?,?)`,
        [data.lenderId, data.streetAddress.value, data.latitude, data.longitude]
      );
      for await (const imageUrl of data.imageUrls) {
        await this._client.query(
          `INSERT INTO images_listing (url, listing_id) VALUES (?,?)`,
          [imageUrl, result.insertId]
        );
      }
      data.id = result.insertId;

      return data;
    } catch (err) {
      this._logger.error(err, "save()");
      throw err;
    }
  }

  public async delete(uid: string): Promise<void> {
    this._logger.info({ uid }, "delete()");
    return await this._client.query(`DELETE FROM listing WHERE id = ?`, [uid]);
  }

  public async findOneById(uid: string): Promise<Listing> {
    this._logger.info({ uid }, "delete()");
    const result = await this._client.query(
      `SELECT * FROM listing WHERE uid = ?`,
      [uid]
    );
    return ListingMapper.toEntityFromRaw(result[0]);
  }

  public async findManyByLatLng(
    latitude: number,
    longitude: number,
    range: number = 100
  ): Promise<Listing[]> {
    this._logger.info({ latitude, longitude, range }, "delete()");
    const result: ListingRawInterface[] = await this._client.query(
      `
        SELECT id, ( 3959 * acos( cos( radians(?) ) * cos( radians( latitude ) ) * cos( radians( longitude ) - radians(?) ) + sin( radians(?) ) * sin( radians( latitude ) ) ) ) 
        AS distance, latitude, longitude 
        FROM listing 
        HAVING distance < ? ORDER BY distance LIMIT 0 , 20
      `,
      [latitude, longitude, latitude, range]
    );

    return result.map((item) => ListingMapper.toEntityFromRaw(item));
  }
}
