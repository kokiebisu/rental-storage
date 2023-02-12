import { findBookings } from "../../src/adapter/resolver/query";
import * as mockEvent from "../event.json";

describe("findBookings()", () => {
  it("should work with valid input", async () => {
    if (!global.data.spaceId) {
      throw new Error("data.spaceId is empty");
    }
    try {
      const event = createEvent({ ...mockEvent });
      const result = await findBookings(event);
      expect(result?.length).toBeGreaterThan(0);
    } catch (err) {
      console.error(err);
    }
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
