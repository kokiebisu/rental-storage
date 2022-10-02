import axios from "axios";

import { BookingPublisherService } from "../../adapter/out/broker/BookingBroker";
import { BookingRepository } from "../port";
import { BookingRepositoryImpl } from "../../adapter/out/db";
import { Amount, Booking } from "../../domain/model";
import { BookingMapper } from "../../adapter/in/mapper";
import { BookingService } from "../port";
import { AWSRegion, Currency } from "../../domain/enum";
import { LoggerUtil } from "../../utils";
import { StorageItemInterface } from "../../types";

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
    guestId: number,
    listingId: number,
    items: StorageItemInterface[]
  ): Promise<boolean> {
    this._logger.info(
      { amount, currency, guestId, listingId },
      "makeBooking()"
    );

    // Check if userId and listId exist
    // code here...
    const { data: user } = await axios.get(
      `${process.env.SERVICE_API_ENDPOINT}/users/guest/${guestId}`
    );
    console.log("USER: ", user);
    if (!user) {
      throw new Error(`Provided userId ${guestId} doesn't exist`);
    }

    const { data: listing } = await axios.get(
      `${process.env.SERVICE_API_ENDPOINT}/listings/${listingId}`
    );
    console.log("LISTING: ", listing);
    if (!listing) {
      throw new Error(`Provided listingId ${listingId} doesn't exist`);
    }

    try {
      const booking = new Booking({
        amount: new Amount({ value: amount, currency }),
        guestId,
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
