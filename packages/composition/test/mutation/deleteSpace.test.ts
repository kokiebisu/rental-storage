import { deleteSpace } from "../../src/adapter/resolver/mutation";
import * as mockEvent from "../event.json";

describe("deleteSpace()", () => {
  it("should work with valid input", async () => {
    if (!global.data.spaceId) {
      throw new Error("data.spaceId is empty");
    }
    const event = createEvent({ ...mockEvent });
    const result = await deleteSpace(event);
    expect(result).not.toBeUndefined();
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
