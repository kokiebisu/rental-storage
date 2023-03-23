import * as mockEvent from "../event.json";
import { acceptBooking } from "../../../src/adapter/resolver/mutation";

describe("acceptBooking()", () => {
  it("should work with valid input", async () => {
    expect(global.data.bookingId).not.toBeUndefined();
    const event = createEvent({ ...mockEvent });
    const result = await acceptBooking(event);
    expect(result.id).not.toBeUndefined();
  });
});

const createEvent = (event: any) => {
  return {
    ...event,
    arguments: {
      bookingId: global.data.bookingId,
    },
    identity: {
      ...event.identity,
      resolverContext: {
        uid: global.data.userId,
      },
    },
  };
};
