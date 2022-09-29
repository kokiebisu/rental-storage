import { AWSRegion, PublisherService } from "@rental-storage-project/common";

export class BookingPublisherService extends PublisherService {
  public static async create(
    region: AWSRegion
  ): Promise<BookingPublisherService> {
    if (!process.env.ACCOUNT_ID) {
      throw new Error("accountId was not successfully retrieved");
    }
    return new BookingPublisherService(region, process.env.ACCOUNT_ID);
  }
}
