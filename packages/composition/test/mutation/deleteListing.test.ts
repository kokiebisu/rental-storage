import { deleteListing } from "../../src/adapter/resolver/mutation";
import * as mockEvent from "../event.json";

describe("deleteListing()", () => {
  it("should work with valid input", async () => {
    if (!global.data.listingId) {
      throw new Error("data.listingId is empty");
    }
    const event = createEvent({ ...mockEvent });
    const result = await deleteListing(event);
    expect(result).not.toBeUndefined();
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
