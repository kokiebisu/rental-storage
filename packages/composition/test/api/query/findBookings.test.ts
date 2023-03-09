import { AppSyncResolverEvent } from "aws-lambda";
import { findBookings } from "../../../src/adapter/resolver/query";
import * as mockEvent from "../event.json";

describe("findBookings()", () => {
  it("should work with valid input", async () => {
    if (!global.data.spaceId) {
      throw new Error("data.spaceId is empty");
    }
    const event = createEvent({ ...mockEvent });
    const result = await findBookings(
      event as AppSyncResolverEvent<{ spaceId: string }, unknown>
    );
    expect(result?.length).toBeGreaterThan(0);
  });
});

const createEvent = (event: AppsyncResolverMockEvent) => {
  return {
    ...event,
    arguments: {
      spaceId: global.data.spaceId,
    },
    identity: {
      ...event.identity,
      resolverContext: {
        uid: global.data.userId,
      },
    },
  };
};
