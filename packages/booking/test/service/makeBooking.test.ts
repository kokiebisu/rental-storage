import { Currency } from "@rental-storage-project/common";
import { BookingServiceImpl } from "../../src/service";

describe("makeBooking()", () => {
  it("shoud work", async () => {
    const service = await BookingServiceImpl.create();
    const data = await service.makeBooking(50, Currency.CAD, "1", "1");
  });
});
