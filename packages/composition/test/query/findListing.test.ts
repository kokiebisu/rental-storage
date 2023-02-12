import { findListing } from "../../src/adapter/resolver/query";
import * as mockEvent from "../event.json";

describe("findListing()", () => {
  it("should work with valid input", async () => {
    if (!global.data.listingId) {
      throw new Error("data.listingId is empty");
    }
    try {
      const event = createEvent({ ...mockEvent });
      const result = await findListing(event);
      expect(result.id).not.toBeUndefined();
    } catch (err) {
      console.error(err);
    }
  });
});

const createEvent = (event: any) => {
  return {
    ...event,
    arguments: {
      id: global.data.listingId,
    },
    identity: {
      ...event.identity,
      resolverContext: {
        uid: global.data.userId,
      },
    },
  };
};
