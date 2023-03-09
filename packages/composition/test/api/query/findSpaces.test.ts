import { AppSyncResolverEvent } from "aws-lambda";
import { findSpaces } from "../../../src/adapter/resolver/query";
import * as mockEvent from "../event.json";

describe("findSpaces()", () => {
  it("should work with valid input", async () => {
    if (!global.data.userId) {
      throw new Error("data.userId is empty");
    }
    const event = createEvent({ ...mockEvent });
    const result = await findSpaces(
      event as AppSyncResolverEvent<{ userId: string }, unknown>
    );
    expect(result.length).toBeGreaterThan(0);
  });
});

const createEvent = (event: AppsyncResolverMockEvent) => {
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
