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
    id: number
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
    const { lenderId, streetAddress, latitude, longitude, items, imageUrls } =
      args;

    const listing = new Listing({
      lenderId,
      streetAddress: new StreetAddress(streetAddress),
      latitude,
      longitude,
      imageUrls,
      items,
    });

    const { data: lender } = await axios.get(
      `${process.env.SERVICE_API_ENDPOINT}/users/${lenderId}`
    );
    if (!lender) {
      throw new Error(`Provided lenderId ${lenderId} doesn't exist`);
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

  public async removeListingById(id: number): Promise<boolean> {
    this._logger.info({ id }, "removeListingById()");
    try {
      await this._listingRepository.delete(id);
      return true;
    } catch (err) {
      this._logger.error(err, "removeListingById()");
      return false;
    }
  }

  public async occupyItem(booking: BookingInterface) {
    this._logger.info({ booking }, "occupyItem()");
    try {
      await this._listingRepository.addItemToListing(
        booking.listingId,
        booking.id
      );
    } catch (err) {
      this._logger.error(err, "occupyItem()");
      throw err;
    }
  }
}
