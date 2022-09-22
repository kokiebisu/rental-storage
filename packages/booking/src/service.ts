import { Booking } from "./entity";
import { BookingRepository } from "./repository/booking";

export interface BookingService {
  makeBooking(): Promise<void>;
  findAllCreatedBookings(): Promise<Booking>;
}

export class BookingServiceImpl {
  private _bookingRepository: BookingRepository;

  private constructor(bookingRepository: Booking) {
    this._bookingRepository = bookingRepository;
  }

  public static async create() {
    const bookingRepository = await BookingRepository.create();

    await bookingRepository.setup();
    return new BookingServiceImpl(bookingRepository);
  }

  public async makeBooking() {}

  public async findAllCreatedBookings() {}
}
