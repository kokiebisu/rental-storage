import { deleteUser } from "../../src/adapter/resolver/mutation";
import * as mockEvent from "../event.json";

describe("deleteUser()", () => {
  it("should work with valid input", async () => {
    if (!global.data.userId) {
      throw new Error("data.userId is empty");
    }
    const event = createEvent({ mockEvent });
    const result = await deleteUser(event);
    expect(result?.id).not.toBeUndefined();
  });
});

const createEvent = (event: any) => {
  return {
    ...event,
    arguments: {
      id: global.data.userId,
    },
    identity: {
      ...event.identity,
      resolverContext: {
        uid: global.data.userId,
      },
    },
  };
};
