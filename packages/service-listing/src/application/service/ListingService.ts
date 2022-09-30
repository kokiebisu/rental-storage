import axios from "axios";

import { ListingMapper } from "../../adapter/in/mapper";
import { ListingRepositoryImpl } from "../../adapter/out/db";
import { Listing, StreetAddress } from "../../domain/model";
import {
  AggregatedListingInterface,
  BookingInterface,
  ListingInterface,
} from "../../types";
import { LoggerUtil } from "../../utils";
import { ListingRepository, ListingService } from "../port";

export class ListingServiceImpl implements ListingService {
  private _listingRepository: ListingRepository;
  private _logger: LoggerUtil;

  private constructor(listingRepository: ListingRepository) {
    this._listingRepository = listingRepository;
    this._logger = new LoggerUtil("ListingServiceImpl");
  }

  public static async create() {
    const listingRepository = await ListingRepositoryImpl.create();

    await listingRepository.setup();
    return new ListingServiceImpl(listingRepository);
  }

  public async findListingsWithinLatLng(
    latitude: number,
    longitude: number,
    range: number
  ): Promise<any> {
    try {
      const listings = await this._listingRepository.findManyByLatLng(
        latitude,
        longitude,
        range
      );
      return listings.map((item) => ListingMapper.toAggregated(item));
    } catch (err) {
      this._logger.error(err, "findListingsWithinLatLng()");
      return null;
    }
  }

  public async findListingById(
    id: string
  ): Promise<AggregatedListingInterface | null> {
    try {
      const listing = await this._listingRepository.findOneById(id);
      return ListingMapper.toAggregated(listing);
    } catch (err) {
      this._logger.error(err, "findListingById()");
      return null;
    }
  }

  public async addListing(
    args: Omit<ListingInterface, "id">
  ): Promise<boolean> {
    this._logger.info(args, "addListing()");
    const { hostId, streetAddress, latitude, longitude, items } = args;

    const listing = new Listing({
      hostId,
      streetAddress: new StreetAddress(streetAddress),
      latitude,
      longitude,
      items,
    });

    const { data: host } = await axios.get(
      `${process.env.SERVICE_API_ENDPOINT}/users/host/${hostId}`
    );
    if (!host) {
      throw new Error(`Provided hostId ${hostId} doesn't exist`);
    }

    try {
      const listingDTO = ListingMapper.toDTOFromEntity(listing);
      await this._listingRepository.save(listingDTO);
      return true;
    } catch (err) {
      this._logger.error(err, "addListing()");
      return false;
    }
  }

  public async removeListingById(id: string): Promise<boolean> {
    try {
      await this._listingRepository.delete(id);
      return true;
    } catch (err) {
      this._logger.error(err, "removeListingById()");
      return false;
    }
  }

  public async occupyStorageItem(booking: BookingInterface) {
    this._logger.info(booking, "occupyStorageItem()");
    await this._listingRepository.addItemToListing(
      booking.listingId,
      booking.items
    );
  }
}
