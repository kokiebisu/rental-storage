import { findListingById } from "../../src/adapter/resolver/query";
import * as mockEvent from "../event.json";

describe("findListingById()", () => {
  it("should work with valid input", async () => {
    if (!global.data.listingId) {
      throw new Error("data.listingId is empty");
    }
    const event = createEvent({ ...mockEvent });
    const result = await findListingById(event);
    expect(result.uid).not.toBeUndefined();
  });
});

const createEvent = (event: any) => {
  return {
    ...event,
    arguments: {
      listingId: global.data.listingId,
    },
    identity: {
      ...event.identity,
      resolverContext: {
        uid: global.data.userId,
      },
    },
  };
};
