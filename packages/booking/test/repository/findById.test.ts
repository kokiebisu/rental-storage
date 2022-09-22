import { AWSRegion } from "@rental-storage-project/common";
import { BookingRepository } from "../../src/repository/booking";

describe.skip("findById()", () => {
  it("shoud work", async () => {
    const repository = await BookingRepository.create(AWSRegion.US_EAST_1);
    const bookingId = "6a892fd2-736e-4e9e-bffc-7511d07770a4";
    const data = await repository.findById(bookingId);
    expect(data.id).not.toBeUndefined();
    expect(data.status).not.toBeUndefined();
    expect(data.amount).not.toBeUndefined();
    expect(data.userId).not.toBeUndefined();
    expect(data.listingId).not.toBeUndefined();
    expect(data.createdAt).not.toBeUndefined();
  });
});
