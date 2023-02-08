import {
  FindBookingByIdCommand,
  FindBookingByIdUseCase,
} from "../../src/adapter/usecase/findBookingById";

describe("findBookingById()", () => {
  it("should work with valid input", async () => {
    if (!global.data.bookingId) {
      throw new Error("data.bookingId is empty");
    }
    const input = { bookingId: global.data.bookingId };
    const usecase = new FindBookingByIdUseCase();
    const result = await usecase.execute(new FindBookingByIdCommand(input));
    expect(result).not.toBeUndefined();
  });
});
