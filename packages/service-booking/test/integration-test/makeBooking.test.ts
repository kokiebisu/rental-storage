import { BookingServiceImpl } from "../../src/app/service";
import { BookingService } from "../../src/app/port";

let service: BookingService;

beforeAll(async () => {
  service = await BookingServiceImpl.create();
});

describe("makeBooking()", () => {
  it("should work", () => {});
});