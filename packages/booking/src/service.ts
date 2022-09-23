import { Amount, AWSRegion, Currency } from "@rental-storage-project/common";

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

  private constructor(
    bookingRepository: BookingRepository,
    bookingPublisher: BookingPublisherService
  ) {
    this._bookingRepository = bookingRepository;
    this._publisher = bookingPublisher;
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
    listingId: string
  ): Promise<boolean> {
    const amountValue: Amount = { value: amount, currency };
    const booking = new Booking({
      amount: amountValue,
      userId,
      listingId,
    });

    try {
      const bookingDTO = BookingMapper.toDTOFromEntity(booking);
      await this._bookingRepository.save(bookingDTO);
      await this._publisher.bookingCreated(bookingDTO);

      return true;
    } catch (err) {
      console.error(err);
      return false;
    }
  }

  public async findAllCreatedBookings() {}
}
