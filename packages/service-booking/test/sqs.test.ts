import { AWSRegion } from "@rental-storage-project/common";
import { BookingPublisherService } from "../src/publisher";

describe.skip("findById()", () => {
  it("shoud work", async () => {
    const publisher = await BookingPublisherService.create(AWSRegion.US_EAST_1);
    await publisher.bookingComplete();
  });
});
