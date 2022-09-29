import {
  Amount,
  AWSRegion,
  Currency,
  LoggerService,
  StorageItemInterface,
} from "@rental-storage-project/common";

import { Booking } from "./entity";
import { BookingMapper } from "./mapper";
import { BookingPublisherService } from "./publisher";
import { BookingRepository } from "./repository/booking";

export interface BookingService {
  makeBooking(): Promise<boolean>;
  findAllCreatedBookings(): Promise<Booking>;
}

export class BookingServiceImpl {
  private _bookingRepository: BookingRepository;
  private _publisher: BookingPublisherService;
  private _logger: LoggerService;

  private constructor(
    bookingRepository: BookingRepository,
    bookingPublisher: BookingPublisherService
  ) {
    this._bookingRepository = bookingRepository;
    this._publisher = bookingPublisher;
    this._logger = new LoggerService("BookingServiceImpl");
  }

  public static async create() {
    const bookingRepository = await BookingRepository.create(
      AWSRegion.US_EAST_1
    );
    const bookingPublisher = await BookingPublisherService.create(
      AWSRegion.US_EAST_1
    );

    return new BookingServiceImpl(bookingRepository, bookingPublisher);
  }

  public async makeBooking(
    amount: number,
    currency: Currency,
    userId: string,
    listingId: string,
    items: StorageItemInterface[]
  ): Promise<boolean> {
    this._logger.info({ amount, currency, userId, listingId }, "makeBooking()");

    try {
      const amountValue: Amount = { value: amount, currency };
      const booking = new Booking({
        amount: amountValue,
        userId,
        listingId,
        items,
      });
      const bookingDTO = BookingMapper.toDTOFromEntity(booking);
      this._logger.debug(bookingDTO, "makeBooking()");
      await this._bookingRepository.save(bookingDTO);
      return true;
    } catch (err) {
      this._logger.error(err, "makeBooking()");
      return false;
    }
  }

  public async findAllCreatedBookings() {}
}
