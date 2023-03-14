import { AppSyncResolverEvent } from "aws-lambda";
import { findBooking } from "../../../src/adapter/resolver/query";
import * as mockEvent from "../event.json";

describe("findBooking()", () => {
  it("should work with valid input", async () => {
    if (!global.data.bookingId) {
      throw new Error("data.bookingId is empty");
    }
    const event = createEvent({ ...mockEvent });
    const result = await findBooking(
      event as AppSyncResolverEvent<{ id: string }, unknown>
    );
    expect(result?.id).not.toBeUndefined();
    expect(result?.spaceId).not.toBeUndefined();

    expect(result?.id.length).toBeGreaterThan(0);
    expect(result?.spaceId.length).toBeGreaterThan(0);
  });
});

const createEvent = (event: AppsyncResolverMockEvent) => {
  return {
    ...event,
    arguments: {
      id: global.data.bookingId,
    },
    identity: {
      ...event.identity,
      resolverContext: {
        uid: global.data.userId,
      },
    },
  };
};
