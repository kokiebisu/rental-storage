import {
  AggregatedListingInterface,
  BookingInterface,
  ListingInterface,
  LoggerService,
} from "@rental-storage-project/common";
import { Listing, StreetAddress } from "./entity";
import { ListingMapper } from "./mapper";
import { ListingRepository } from "./repository";

interface ListingService {
  findListingsWithinLatLng(
    latitude: number,
    longitude: number,
    range: number
  ): Promise<AggregatedListingInterface[]>;
  findListingById(id: string): Promise<AggregatedListingInterface | null>;
  addListing(data: ListingInterface): Promise<boolean>;
  removeListingById(id: string): Promise<boolean>;
}

export class ListingServiceImpl implements ListingService {
  private _listingRepository: ListingRepository;
  private _logger: LoggerService;

  private constructor(listingRepository: ListingRepository) {
    this._listingRepository = listingRepository;
    this._logger = new LoggerService("ListingServiceImpl");
  }

  public static async create() {
    const listingRepository = await ListingRepository.create();

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
    const { hostId, streetAddress, latitude, longitude, items } = args;
    const listing = new Listing(
      hostId,
      new StreetAddress(streetAddress),
      latitude,
      longitude,
      items
    );
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
    await this._listingRepository.addItemToListing(booking.id, booking.items);
  }
}
