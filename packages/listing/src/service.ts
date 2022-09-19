import {
  AggregatedListingInterface,
  Listing,
  ListingInterface,
  StreetAddress,
} from "./entity";
import { EmailAddress } from "./entity/email-address";
import { ListingMapper } from "./mapper";
import { ListingRepository } from "./repository";

interface ListingService {
  findListingById(id: string): Promise<AggregatedListingInterface | null>;
  addListing(data: ListingInterface): Promise<boolean>;
  removeListingById(id: string): Promise<boolean>;
}

export class ListingServiceImpl implements ListingService {
  private _listingRepository: ListingRepository;

  private constructor(listingRepository: ListingRepository) {
    this._listingRepository = listingRepository;
  }

  public static async create() {
    const listingRepository = await ListingRepository.create();

    await listingRepository.setup();
    return new ListingServiceImpl(listingRepository);
  }

  public async findListingById(
    id: string
  ): Promise<AggregatedListingInterface | null> {
    try {
      const listing = await this._listingRepository.findOneById(id);
      const aggregatedListing = ListingMapper.toAggregatedDTO(listing);
      return aggregatedListing;
    } catch (err) {
      console.error(err);
      return null;
    }
  }

  public async addListing(
    args: Omit<ListingInterface, "id">
  ): Promise<boolean> {
    const { hostId, emailAddress, streetAddress, latitude, longitude } = args;
    const listing = new Listing(
      hostId,
      new EmailAddress(emailAddress),
      new StreetAddress(streetAddress),
      latitude,
      longitude
    );
    try {
      await this._listingRepository.save(
        ListingMapper.toDTOFromEntity(listing)
      );
      return true;
    } catch (err) {
      console.error(err);
      return false;
    }
  }

  public async removeListingById(id: string): Promise<boolean> {
    try {
      await this._listingRepository.delete(id);
      return true;
    } catch (err) {
      console.error(err);
      return false;
    }
  }
}
