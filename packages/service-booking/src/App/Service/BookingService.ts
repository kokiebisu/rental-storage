import axios from "axios";

import { BookingPublisherService } from "../../Adapter/Broker/BookingBroker";
import { BookingRepository } from "../port";
import { BookingRepositoryImpl } from "../../adapter/database";
import { Amount, Booking } from "../../Domain/Model";
import { BookingMapper } from "../../adapter/mapper";
import { BookingService } from "../port";
import { AWSRegion } from "../../Domain/Enum";
import { LoggerUtil } from "../../Utils";
import { ItemInterface } from "../../Types";

export class BookingServiceImpl implements BookingService {
  private _bookingRepository: BookingRepository;
  private _publisher: BookingPublisherService;
  private _logger: LoggerUtil;

  private constructor(
    bookingRepository: BookingRepository,
    bookingPublisher: BookingPublisherService
  ) {
    this._bookingRepository = bookingRepository;
    this._publisher = bookingPublisher;
    this._logger = new LoggerUtil("BookingServiceImpl");
  }

  public static async create() {
    const bookingRepository = await BookingRepositoryImpl.create(
      AWSRegion.US_EAST_1
    );
    const bookingPublisher = await BookingPublisherService.create(
      AWSRegion.US_EAST_1
    );

    return new BookingServiceImpl(bookingRepository, bookingPublisher);
  }

  public async makeBooking(
    amount: number,
    currency: string,
    ownerId: number,
    listingId: number,
    items: ItemInterface[]
  ): Promise<boolean> {
    this._logger.info(
      { amount, currency, ownerId, listingId },
      "makeBooking()"
    );

    // Check if userId and listId exist
    // code here...
    const { data: user } = await axios.get(
      `${process.env.SERVICE_API_ENDPOINT}/users/${ownerId}`
    );
    if (!user) {
      throw new Error(`Provided userId ${ownerId} doesn't exist`);
    }

    const { data: listing } = await axios.get(
      `${process.env.SERVICE_API_ENDPOINT}/listings/${listingId}`
    );
    if (!listing) {
      throw new Error(`Provided listingId ${listingId} doesn't exist`);
    }

    try {
      const booking = new Booking({
        amount: new Amount({ value: amount, currency }),
        ownerId,
        listingId,
        items,
      });

      await this._bookingRepository.save(
        BookingMapper.toDTOFromEntity(booking)
      );
      return true;
    } catch (err) {
      this._logger.error(err, "makeBooking()");
      return false;
    }
  }

  public async findAllCreatedBookings() {
    return true;
  }
}
