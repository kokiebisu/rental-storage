import { AppSyncResolverEvent } from "aws-lambda";
import { findBookings } from "../../../src/adapter/resolver/query";
import * as mockEvent from "../event.json";

describe("findBookings()", () => {
  describe('when the bookingStatus is "pending"', () => {
    it("should work with valid input", async () => {
      if (!global.data.spaceId) {
        throw new Error("data.spaceId is empty");
      }
      const event = createEvent({ ...mockEvent }, "pending");
      const result = await findBookings(
        event as AppSyncResolverEvent<
          { spaceId: string; bookingStatus: "pending" | "approved" },
          unknown
        >
      );
      expect(result?.length).toBeGreaterThan(0);
    });
  });

  describe('when the bookingStatus is "approved"', () => {
    it("should work with valid input", async () => {
      if (!global.data.spaceId) {
        throw new Error("data.spaceId is empty");
      }
      const event = createEvent({ ...mockEvent }, "approved");
      const result = await findBookings(
        event as AppSyncResolverEvent<
          { spaceId: string; bookingStatus: "pending" | "approved" },
          unknown
        >
      );
      expect(result?.length).toBe(0);
    });
  });
});

const createEvent = (
  event: AppsyncResolverMockEvent,
  bookingStatus: "pending" | "approved"
) => {
  return {
    ...event,
    arguments: {
      spaceId: global.data.spaceId,
      bookingStatus: bookingStatus,
    },
    identity: {
      ...event.identity,
      resolverContext: {
        uid: global.data.userId,
      },
    },
  };
};
