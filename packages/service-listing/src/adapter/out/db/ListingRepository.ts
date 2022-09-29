import Client from "serverless-mysql";
import {
  ListingInterface,
  LoggerService,
  StorageItemInterface,
} from "@rental-storage-project/common";
import { ListingMapper } from "../../in/mapper";
import { ListingRepository } from "../../../application/port";
import { ListingRawInterface } from "../../../domain/model";

export class ListingRepositoryImpl implements ListingRepository {
  public readonly tableName: string;
  private _client: any;
  private _logger: LoggerService;

  private constructor(client: any, tableName: string, className: string) {
    this._client = client;
    this.tableName = tableName;
    this._logger = new LoggerService(className);
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
      "CREATE TABLE IF NOT EXISTS listing (id INT AUTO_INCREMENT NOT NULL, host_id INT, street_address VARCHAR(100) NOT NULL, latitude INT NOT NULL, longitude INT NOT NULL, PRIMARY KEY (id))"
    );
    await this._client.query(
      "CREATE TABLE IF NOT EXISTS listing_item (listing_id INT NOT NULL,item_id INT NOT NULL, PRIMARY KEY (listing_id, item_id))"
    );
  }

  public async save(data: Omit<ListingInterface, "id">): Promise<void> {
    await this._client.query(
      `INSERT INTO listing (host_id, street_address, latitude, longitude) VALUES(?,?,?,?)`,
      [data.hostId, data.streetAddress, data.latitude, data.longitude]
    );
  }

  public async addItemToListing(
    listingId: string,
    itemId: StorageItemInterface[]
  ) {
    this._logger.info({ listingId, itemId }, "addItemToListing()");

    await this._client.query(
      `INSERT INTO listing_item (listing_id, item_id) VALUES (?,?)`,
      listingId,
      itemId
    );
  }

  public async delete(id: string): Promise<void> {
    return await this._client.query(`DELETE FROM listing WHERE id = ?`, [id]);
  }

  public async findOneById(listingId: string): Promise<ListingInterface> {
    const result = await this._client.query(
      `SELECT * FROM listing WHERE id = ?`,
      [listingId]
    );
    return ListingMapper.toDTOFromRaw(result[0]);
  }

  public async findManyByLatLng(
    latitude: number,
    longitude: number,
    range: number = 100
  ): Promise<ListingInterface[]> {
    const result: ListingRawInterface[] = await this._client.query(
      `
      SELECT id, ( 3959 * acos( cos( radians(?) ) * cos( radians( latitude ) ) * cos( radians( longitude ) - radians(?) ) + sin( radians(?) ) * sin( radians( latitude ) ) ) ) AS distance, latitude, longitude FROM listing HAVING distance < ? ORDER BY distance LIMIT 0 , 20
    `,
      [latitude, longitude, latitude, range]
    );

    return result.map((item) => ListingMapper.toDTOFromRaw(item));
  }
}
