import {
  AWSRegion,
  BookingInterface,
  PublisherService,
} from "@rental-storage-project/common";

export class BookingPublisherService extends PublisherService {
  public async create(region: AWSRegion): Promise<BookingPublisherService> {
    return new BookingPublisherService(region);
  }

  public async bookingCreated(booking: BookingInterface) {
    const topicArn = `${process.env.namespace}-${process.env.service}-${process.env.stage}-bookingCreated`;
    const message = JSON.parse(booking);
    await this.publish(message, topicArn);
  }
}
