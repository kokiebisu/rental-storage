import axios from "axios";

import { ListingMapper } from "../../adapter/in/mapper";
import { ListingRepositoryImpl } from "../../adapter/out/db";
import { Listing, StreetAddress } from "../../domain/model";
import { ListingInterface } from "../../types";
import { LoggerUtil } from "../../utils";
import { AddListing, ListingRepository, ListingService } from "../port";

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
  ): Promise<ListingInterface[]> {
    this._logger.info(
      { latitude, longitude, range },
      "findListingsWithinLatLng()"
    );
    try {
      const listings = await this._listingRepository.findManyByLatLng(
        latitude,
        longitude,
        range
      );
      return listings.map((item) => ListingMapper.toDTOFromEntity(item));
    } catch (err) {
      this._logger.error(err, "findListingsWithinLatLng()");
      throw err;
    }
  }

  public async findListingById(uid: string): Promise<ListingInterface> {
    this._logger.info({ uid }, "findListingById()");
    try {
      const listing = await this._listingRepository.findOneById(uid);

      // return ListingMapper.toAggregated(listing);
      return ListingMapper.toDTOFromEntity(listing);
    } catch (err) {
      this._logger.error(err, "findListingById()");
      throw err;
    }
  }

  public async addListing(args: AddListing): Promise<boolean> {
    this._logger.info(args, "addListing()");
    const { lenderId, streetAddress, latitude, longitude, imageUrls } = args;
    try {
      const listing = new Listing({
        lenderId,
        streetAddress: new StreetAddress(streetAddress),
        latitude,
        longitude,
        imageUrls,
      });

      if (!this.checkListingExists(lenderId)) {
        throw new Error(`Provided lenderId ${lenderId} doesn't exist`);
      }

      await this._listingRepository.save(listing);
      return true;
    } catch (err) {
      this._logger.error(err, "addListing()");
      return false;
    }
  }

  public async removeListingById(uid: string): Promise<boolean> {
    this._logger.info({ uid }, "removeListingById()");
    try {
      const listing = await this._listingRepository.findOneById(uid);
      await this._listingRepository.delete(listing);
      return true;
    } catch (err) {
      this._logger.error(err, "removeListingById()");
      return false;
    }
  }

  // public async occupyStorageItem(booking: BookingInterface) {
  //   this._logger.info({ booking }, "occupyStorageItem()");
  //   try {
  //     await this._listingRepository.addItemToListing(
  //       booking.listingId,
  //       booking.id
  //     );
  //   } catch (err) {
  //     this._logger.error(err, "occupyStorageItem()");
  //     throw err;
  //   }
  // }

  private async checkListingExists(lenderId: string) {
    this._logger.info({ lenderId }, "checkListingExists()");
    // check if len
    const { data: lender } = await axios.get(
      `${process.env.SERVICE_API_ENDPOINT}/users/${lenderId}`
    );
    return lender;
  }
}
