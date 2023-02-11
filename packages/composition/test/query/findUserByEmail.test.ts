import { findUserByEmail } from "../../src/adapter/resolver/query";
import * as mockEvent from "../event.json";

describe("findUserByEmail()", () => {
  it("should work with valid input", async () => {
    if (!global.data.mockEmailAddress) {
      throw new Error("data.mockEmailAddress is empty");
    }
    const event = createEvent({ ...mockEvent });
    const result = await findUserByEmail(event);
    expect(result.uid).not.toBeUndefined();
  });
});

const createEvent = (event: any) => {
  return {
    ...event,
    arguments: {
      emailAddress: global.data.mockEmailAddress,
    },
    identity: {
      ...event.identity,
      resolverContext: {
        uid: global.data.userId,
      },
    },
  };
};
