import { Client } from "pg";

import { ListingMapper } from "../mapper";
import { ListingRepository } from "../../application/port";
import { LoggerUtil } from "../../utils";
import { Listing } from "../../domain/model";

export class ListingRepositoryImpl implements ListingRepository {
  public readonly tableName: string;
  private _logger: LoggerUtil;

  private constructor(tableName: string, className: string) {
    this.tableName = tableName;
    this._logger = new LoggerUtil(className);
  }
  public static async create(): Promise<ListingRepositoryImpl> {
    return new ListingRepositoryImpl("listing", "ListingRepository");
  }

  public getDBClient(): Client {
    const client = new Client({
      host: process.env.DB_HOST,
      database: process.env.DB_NAME,
      user: process.env.DB_USERNAME,
      password: process.env.DB_PASSWORD,
      port: Number(process.env.DB_PORT),
    });

    return client;
  }

  public async setup(): Promise<void> {
    this._logger.info({}, "setup()");
    const client = this.getDBClient();
    try {
      await client.connect();
      await client.query(
        `
          CREATE TABLE IF NOT EXISTS listing (
            id SERIAL NOT NULL PRIMARY KEY, 
            uid VARCHAR(64) NOT NULL,
            lender_id VARCHAR(64), 
            street_address VARCHAR(100) NOT NULL, 
            latitude DECIMAL(16,14) NOT NULL, 
            longitude DECIMAL(16,14) NOT NULL
          )
        `
      );
      await client.query(
        `
          CREATE TABLE IF NOT EXISTS images_listing (
            id SERIAL NOT NULL PRIMARY KEY,
            url TEXT,
            listing_id INT NOT NULL,
            CONSTRAINT fk_listing
              FOREIGN KEY(listing_id) 
                REFERENCES listing(id)
                ON DELETE CASCADE
          )
        `
      );
      await client.end();
    } catch (err) {
      this._logger.error(err, "setup()");
      await client.end();
      throw err;
    }
  }

  public async save(data: Listing): Promise<Listing> {
    this._logger.info(data, "save()");
    const client = this.getDBClient();
    try {
      await client.connect();
      // commit/transaction add rollback behavior
      const result = await client.query(
        `
          INSERT INTO listing (
          uid, lender_id, street_address, latitude, longitude
          ) VALUES ($1, $2, $3, $4, $5)
          RETURNING *`,
        [
          data.uid,
          data.lenderId,
          data.streetAddress.value,
          data.latitude,
          data.longitude,
        ]
      );
      for await (const imageUrl of data.imageUrls) {
        await client.query(
          `INSERT INTO images_listing (url, listing_id) VALUES ($1, $2)`,
          [imageUrl, result.rows[0].id]
        );
      }
      data.id = result.rows[0].id;
      await client.end();
      return data;
    } catch (err) {
      this._logger.error(err, "save()");
      await client.end();
      throw err;
    }
  }

  public async delete(data: Listing): Promise<void> {
    this._logger.info({ data }, "delete()");
    const client = this.getDBClient();
    // commit/transaction remove listing and associated urls
    try {
      await client.connect();
      const result = await client.query(
        `DELETE FROM listing WHERE uid = $1 RETURNING *`,
        [data.uid]
      );
      await client.query(
        `DELETE FROM images_listing WHERE listing_id = $1 RETURNING *`,
        [result.rows[0].id]
      );
    } catch (err) {
      this._logger.error(err, "delete()");
      await client.end();
      throw err;
    }
  }

  public async findOneById(uid: string): Promise<Listing> {
    this._logger.info({ uid }, "findOneById()");
    const client = this.getDBClient();
    try {
      await client.connect();
      const result = await client.query(
        `
          SELECT * FROM listing 
          LEFT JOIN images_listing ON listing.id = images_listing.listing_id
          WHERE listing.uid = $1
        `,
        [uid]
      );
      await client.end();
      return ListingMapper.toEntityFromRaw(result.rows[0]);
    } catch (err) {
      this._logger.error(err, "findOneById()");
      await client.end();
      throw err;
    }
  }

  public async findManyByUserId(userId: string): Promise<Listing[]> {
    this._logger.info({ userId }, "findManyByUserId()");
    const client = this.getDBClient();
    try {
      await client.connect();
      const result = await client.query(
        `
          SELECT * FROM listing 
          LEFT JOIN images_listing ON listing.id = images_listing.listing_id
          WHERE listing.lender_id = $1
        `,
        [userId]
      );
      await client.end();
      return result.rows.map((item) => ListingMapper.toEntityFromRaw(item));
    } catch (err) {
      this._logger.error(err, "findManyByUserId()");
      await client.end();
      throw err;
    }
  }

  public async findManyByLatLng(
    latitude: number,
    longitude: number,
    range: number = 100
  ): Promise<Listing[]> {
    this._logger.info({ latitude, longitude, range }, "findManyByLatLng()");
    const client = this.getDBClient();
    try {
      await client.connect();
      const result = await client.query(
        `
          SELECT listing.*, ( 3959 * acos( cos( radians($1) ) * cos( radians( latitude ) ) * cos( radians( longitude ) - radians($2) ) + sin( radians($3) ) * sin( radians( latitude ) ) ) ) 
          AS distance, images_listing.url
          FROM listing INNER JOIN images_listing ON listing.id = images_listing.id
          HAVING distance < $4 ORDER BY distance LIMIT 0 , 20
        `,
        [latitude, longitude, latitude, range]
      );
      await client.end();
      return result.rows.map((item) => ListingMapper.toEntityFromRaw(item));
    } catch (err) {
      this._logger.error(err, "findManyByLatLng()");
      await client.end();
      throw err;
    }
  }
}
