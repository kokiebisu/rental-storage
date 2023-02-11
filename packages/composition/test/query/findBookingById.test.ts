import { findBookingById } from "../../src/adapter/resolver/query";
import * as mockEvent from "../event.json";

describe("findBookingById()", () => {
  it("should work with valid input", async () => {
    if (!global.data.bookingId) {
      throw new Error("data.bookingId is empty");
    }
    const event = createEvent({ ...mockEvent });
    const result = await findBookingById(event);
    expect(result?.uid).not.toBeUndefined();
    expect(result?.items).not.toBeUndefined();
    expect(result?.listingId).not.toBeUndefined();
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
