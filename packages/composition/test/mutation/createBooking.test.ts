import { mock } from "../mock";
import * as mockEvent from "../event.json";
import { createBooking } from "../../src/adapter/resolver/mutation";

describe("createBooking()", () => {
  it("should work with valid input", async () => {
    expect(global.data.bookingId).not.toBeUndefined();
    const event = createEvent({ ...mockEvent });
    const result = await createBooking(event);
    expect(result.uid).not.toBeUndefined();
  });
});

const createEvent = (event: any) => {
  return {
    ...event,
    arguments: {
      spaceId: global.data.spaceId,
    },
    identity: {
      ...event.identity,
      resolverContext: {
        uid: global.data.userId,
      },
    },
  };
};
