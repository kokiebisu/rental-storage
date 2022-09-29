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

  public async findAllCreatedBookings() {
    return true;
  }
}
