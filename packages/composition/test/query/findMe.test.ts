import { findMe } from "../../src/adapter/resolver/query";
import * as mockEvent from "../event.json";

describe("findMe()", () => {
  it("should work with valid input", async () => {
    if (!global.data.userId) {
      throw new Error("data.uid is empty");
    }
    const event = createEvent({ ...mockEvent });
    const result = await findMe(event);
    expect(result.uid).not.toBeUndefined();
  });
});

const createEvent = (event: any) => {
  return {
    ...event,
    arguments: {},
    identity: {
      ...event.identity,
      resolverContext: {
        uid: global.data.userId,
      },
    },
  };
};
