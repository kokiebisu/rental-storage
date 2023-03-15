import { AppSyncResolverEvent } from "aws-lambda";
import { AxiosError } from "axios";
import { findSpaces } from "../../../src/adapter/resolver/query";
import * as mockEvent from "../event.json";

describe("findSpaces()", () => {
  it("shouldn't work without any filters", async () => {
    if (!global.data.userId) {
      throw new Error("data.userId is empty");
    }
    const event = createEvent({ ...mockEvent }, {});
    expect(() =>
      findSpaces(
        event as AppSyncResolverEvent<{ filter?: { userId?: string } }, unknown>
      )
    ).rejects.toThrowError(AxiosError);
  });

  it('should work with "offset" and "limit" as filter', async () => {
    if (!global.data.userId) {
      throw new Error("data.userId is empty");
    }
    const event = createEvent(
      { ...mockEvent },
      {
        filter: {
          offset: 0,
          limit: 10,
        },
      }
    );
    const result = await findSpaces(
      event as AppSyncResolverEvent<
        { filter?: { userId?: string; offset?: number; limit?: number } },
        unknown
      >
    );
    expect(result.length).toBeGreaterThan(0);
  });

  it('should work with "userId" as filter', async () => {
    if (!global.data.userId) {
      throw new Error("data.userId is empty");
    }
    const event = createEvent(
      { ...mockEvent },
      {
        filter: {
          userId: global.data.userId,
        },
      }
    );
    const result = await findSpaces(
      event as AppSyncResolverEvent<
        { filter?: { userId?: string; offset?: number; limit?: number } },
        unknown
      >
    );
    expect(result.length).toBeGreaterThan(0);
  });
});

const createEvent = (
  event: AppsyncResolverMockEvent,
  args: {
    filter?: { userId?: string; offset?: number; limit?: number };
  } = {}
) => {
  return {
    ...event,
    arguments: args,
    identity: {
      ...event.identity,
      resolverContext: {
        uid: global.data.userId,
      },
    },
  };
};
