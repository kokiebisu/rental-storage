import { mock } from "../mock";
import * as mockEvent from "../event.json";
import { makeBooking } from "../../src/adapter/resolver/mutation";

describe("makeBooking()", () => {
  it("should work with valid input", async () => {
    expect(global.data.bookingId).not.toBeUndefined();
    const event = createEvent({ ...mockEvent });
    const result = await makeBooking(event);
    expect(result.uid).not.toBeUndefined();
  });
});

const createEvent = (event: any) => {
  return {
    ...event,
    arguments: {
      listingId: global.data.listingId,
      items: mock.items,
    },
    identity: {
      ...event.identity,
      resolverContext: {
        uid: global.data.userId,
      },
    },
  };
};
