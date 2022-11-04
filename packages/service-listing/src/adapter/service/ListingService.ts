import axios from "axios";

import { ListingMapper } from "../mapper";
import { ListingRepositoryImpl } from "../repository";
import { ListingInterface } from "../../domain/types";
import { LoggerUtil } from "../../utils";
import { AddListing, ListingService } from "../../port/service";
import { ListingKinesisStreamEventSender } from "../sender-event";
import { AWSRegion, CurrencyType, RentalFeeType } from "../../domain/enum";
import { Amount, Fee, Listing, StreetAddress } from "../../domain/model";
import { ListingRepository } from "../../port/repository";
import { ListingEventSender } from "../../port/sender-event";

export class ListingServiceImpl implements ListingService {
  private _listingRepository: ListingRepository;
  private _listingEventSender: ListingEventSender;
  private _logger: LoggerUtil;

  private constructor(
    listingRepository: ListingRepository,
    listingEventSender: ListingEventSender
  ) {
    this._listingRepository = listingRepository;
    this._listingEventSender = listingEventSender;
    this._logger = new LoggerUtil("ListingServiceImpl");
  }

  public static async create() {
    const listingRepository = await ListingRepositoryImpl.create();
    const listingEventSender = await ListingKinesisStreamEventSender.create(
      AWSRegion.US_EAST_1
    );

    await listingRepository.setup();
    return new ListingServiceImpl(listingRepository, listingEventSender);
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

  public async findListingsByUserId(
    userId: string
  ): Promise<ListingInterface[]> {
    this._logger.info({ userId }, "findListingsByUserId()");
    try {
      const listings = await this._listingRepository.findManyByUserId(userId);

      // return ListingMapper.toAggregated(listing);
      return listings.map((listing) => ListingMapper.toDTOFromEntity(listing));
    } catch (err) {
      this._logger.error(err, "findListingById()");
      throw err;
    }
  }

  public async addListing(args: AddListing): Promise<boolean> {
    this._logger.info(args, "addListing()");
    const {
      lenderId,
      streetAddress,
      latitude,
      longitude,
      imageUrls,
      title,
      fee,
    } = args;
    console.log("FEE: ", fee.amount, fee.currency, fee.type);
    try {
      let listing = new Listing({
        title,
        lenderId,
        streetAddress: new StreetAddress(streetAddress),
        latitude,
        longitude,
        imageUrls,
        fee: new Fee({
          amount: new Amount({
            value: fee.amount,
            currency: fee.currency as CurrencyType,
          }),
          type: fee.type as RentalFeeType,
        }),
      });

      if (!this.checkListingExists(lenderId)) {
        throw new Error(`Provided lenderId ${lenderId} doesn't exist`);
      }

      listing = await this._listingRepository.save(listing);
      const listingDTO = ListingMapper.toDTOFromEntity(listing);
      await this._listingEventSender.listingCreated(listingDTO);
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

  private async checkListingExists(lenderId: string): Promise<any> {
    this._logger.info({ lenderId }, "checkListingExists()");
    // check if len
    const { data: lender } = await axios.get(
      `${process.env.SERVICE_API_ENDPOINT}/users/${lenderId}`
    );
    return lender;
  }
}
