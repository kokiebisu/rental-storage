import { findListings } from "../../src/adapter/resolver/query";
import * as mockEvent from "../event.json";

describe("findListings()", () => {
  it("should work with valid input", async () => {
    if (!global.data.userId) {
      throw new Error("data.userId is empty");
    }
    const event = createEvent({ ...mockEvent });
    const result = await findListings(event);
    expect(result.listings.length).toBeGreaterThan(0);
  });
});

const createEvent = (event: any) => {
  return {
    ...event,
    arguments: {
      userId: global.data.userId,
    },
    identity: {
      ...event.identity,
      resolverContext: {
        uid: global.data.userId,
      },
    },
  };
};
