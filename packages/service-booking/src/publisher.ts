import {
  AWSRegion,
  BookingInterface,
  PublisherService,
} from "@rental-storage-project/common";

export class BookingPublisherService extends PublisherService {
  public static async create(
    region: AWSRegion
  ): Promise<BookingPublisherService> {
    if (!process.env.account_id) {
      throw new Error("account_id was not successfully retrieved");
    }
    return new BookingPublisherService(region, process.env.account_id);
  }

  public async bookingCreated(booking: BookingInterface) {
    const topicArn = `${process.env.namespace}-${process.env.service}-${process.env.stage}-bookingCreated`;
    const message = JSON.stringify(booking);
    try {
      await this.publish(message, topicArn);
      this._logger.info(topicArn, "bookingCreated()");
    } catch (err) {
      this._logger.error(err, "bookingCreated()");
    }
  }
}
