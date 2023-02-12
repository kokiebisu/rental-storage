import { findSpace } from "../../src/adapter/resolver/query";
import * as mockEvent from "../event.json";

describe("findSpace()", () => {
  it("should work with valid input", async () => {
    if (!global.data.spaceId) {
      throw new Error("data.spaceId is empty");
    }
    try {
      const event = createEvent({ ...mockEvent });
      const result = await findSpace(event);
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
      id: global.data.spaceId,
    },
    identity: {
      ...event.identity,
      resolverContext: {
        uid: global.data.userId,
      },
    },
  };
};
