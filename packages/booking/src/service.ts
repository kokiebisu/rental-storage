import { Amount, AWSRegion, Currency } from "@rental-storage-project/common";

import { Booking } from "./entity";
import { BookingMapper } from "./mapper";
import { BookingRepository } from "./repository/booking";

export interface BookingService {
  makeBooking(): Promise<boolean>;
  findAllCreatedBookings(): Promise<Booking>;
}

export class BookingServiceImpl {
  private _bookingRepository: BookingRepository;

  private constructor(bookingRepository: BookingRepository) {
    this._bookingRepository = bookingRepository;
  }

  public static async create() {
    const bookingRepository = await BookingRepository.create(
      AWSRegion.US_EAST_1
    );

    return new BookingServiceImpl(bookingRepository);
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
      await this._bookingRepository.save(
        BookingMapper.toDTOFromEntity(booking)
      );
      return true;
    } catch (err) {
      console.error(err);
      return false;
    }
  }

  public async findAllCreatedBookings() {}
}
