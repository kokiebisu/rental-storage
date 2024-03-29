import { AppSyncResolverEvent } from "aws-lambda";
import { findUser } from "../../../src/adapter/resolver/query";
import * as mockEvent from "../event.json";

describe("findUser()", () => {
  it("should work with valid input", async () => {
    if (!global.data.userId) {
      throw new Error("data.userId is empty");
    }
    const event = createEvent({ ...mockEvent });
    const result = await findUser(
      event as AppSyncResolverEvent<{ id: string }, unknown>
    );
    expect(result).not.toBeUndefined();
  });
});

const createEvent = (event: AppsyncResolverMockEvent) => {
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
