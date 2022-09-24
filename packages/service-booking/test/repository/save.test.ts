import { AWSRegion, Currency } from "@rental-storage-project/common";
import { Booking } from "../../src/entity";
import { BookingMapper } from "../../src/mapper";
import { BookingRepository } from "../../src/repository/booking";

describe.skip("save()", () => {
  it("shoud work", async () => {
    const repository = await BookingRepository.create(AWSRegion.US_EAST_1);
    const booking = new Booking({
      amount: {
        value: 50,
        currency: Currency.CAD,
      },
      userId: "12",
      listingId: "12",
    });
    await repository.save(BookingMapper.toDTOFromEntity(booking));
  });
});
