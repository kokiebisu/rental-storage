import { Client } from "pg";

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
    var client = new Client({
      host: process.env.DB_HOST,
      database: process.env.DB_NAME,
      user: process.env.DB_USERNAME,
      password: process.env.DB_PASSWORD,
      port: Number(process.env.DB_PORT),
    });
    return new ListingRepositoryImpl(client, "listing", "ListingRepository");
  }

  public async setup(): Promise<void> {
    await this._client.query(
      `
        CREATE TABLE IF NOT EXISTS listing (
          id SERIAL NOT NULL PRIMARY KEY, 
          uid CHAR(32) NOT NULL,
          lender_id CHAR(32), 
          street_address VARCHAR(100) NOT NULL, 
          latitude DECIMAL(11,7) NOT NULL, 
          longitude DECIMAL(11,7) NOT NULL
        )
      `
    );
    await this._client.query(
      `
        CREATE TABLE IF NOT EXISTS images_listing (
          id SERIAL NOT NULL PRIMARY KEY,
          url VARCHAR(100),
          listing_id INT NOT NULL,
          CONSTRAINT fk_listing
            FOREIGN KEY(listing_id) 
	            REFERENCES listing(id)
	            ON DELETE CASCADE
        )
      `
    );
  }

  public async save(data: Listing): Promise<Listing> {
    this._logger.info(data, "save()");
    try {
      // commit/transaction add rollback behavior
      const result = await this._client.query(
        `INSERT INTO listing (
          uid, lender_id, street_address, latitude, longitude
        ) VALUES ($1, $2, $3, $4, $5)`,
        [
          data.uid,
          data.lenderId,
          data.streetAddress.value,
          data.latitude,
          data.longitude,
        ]
      );
      for await (const imageUrl of data.imageUrls) {
        await this._client.query(
          `INSERT INTO images_listing (url, listing_id) VALUES ($1, $2)`,
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

  public async delete(data: Listing): Promise<void> {
    this._logger.info({ data }, "delete()");
    // commit/transaction remove listing and associated urls
    try {
      await this._client.query(`DELETE FROM listing WHERE uid = $1`, [
        data.uid,
      ]);
      await this._client.query(
        `DELETE FROM images_listing WHERE listing_id = $1`,
        [data.id]
      );
    } catch (err) {
      this._logger.error(err, "delete()");
      throw err;
    }
  }

  public async findOneById(uid: string): Promise<Listing> {
    const result = await this._client.query(
      `
        SELECT * FROM listing 
        INNER JOIN images_listing ON listing.id = images_listing.id
        WHERE listing.uid = $1
      `,
      [uid]
    );

    return ListingMapper.toEntityFromRaw(result[0]);
  }

  public async findManyByLatLng(
    latitude: number,
    longitude: number,
    range: number = 100
  ): Promise<Listing[]> {
    this._logger.info({ latitude, longitude, range }, "findManyByLatLng()");
    const result: ListingRawInterface[] = await this._client.query(
      `
        SELECT listing.*, ( 3959 * acos( cos( radians($1) ) * cos( radians( latitude ) ) * cos( radians( longitude ) - radians($2) ) + sin( radians($3) ) * sin( radians( latitude ) ) ) ) 
        AS distance, images_listing.url
        FROM listing INNER JOIN images_listing ON listing.id = images_listing.id
        HAVING distance < $4 ORDER BY distance LIMIT 0 , 20
      `,
      [latitude, longitude, latitude, range]
    );

    return result.map((item) => ListingMapper.toEntityFromRaw(item));
  }
}
