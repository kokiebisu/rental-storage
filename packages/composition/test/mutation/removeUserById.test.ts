import { removeUserById } from "../../src/adapter/resolver/mutation";
import * as mockEvent from "../event.json";

describe("removeUserById()", () => {
  it("should work with valid input", async () => {
    if (!global.data.userId) {
      throw new Error("data.userId is empty");
    }
    const event = createEvent({ mockEvent });
    const result = await removeUserById(event);
    expect(result?.uid).not.toBeUndefined();
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
