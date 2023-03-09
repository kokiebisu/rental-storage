import { AppSyncResolverEvent } from "aws-lambda";
import { findSpace } from "../../../src/adapter/resolver/query";
import * as mockEvent from "../event.json";

describe("findSpace()", () => {
  it("should work with valid input", async () => {
    if (!global.data.spaceId) {
      throw new Error("data.spaceId is empty");
    }

    const event = createEvent({ ...mockEvent });
    const result = await findSpace(
      event as AppSyncResolverEvent<{ id: string }, unknown>
    );
    expect(result.id).not.toBeUndefined();
    expect(result.lenderId).not.toBeUndefined();
    expect(result.location).not.toBeUndefined();
    expect(result.imageUrls.length).toBeGreaterThan(0);
    expect(result.title).not.toBeUndefined();
    expect(result.description).not.toBeUndefined();
    expect(result.createdAt).not.toBeUndefined();
    expect(result.updatedAt).not.toBeUndefined();
  });
});

const createEvent = (event: AppsyncResolverMockEvent) => {
  return {
    ...event,
    arguments: {
      id: global.data.spaceId,
    },
    identity: {
      ...event.identity,
      resolverContext: {
        uid: global.data.userId,
      },
    },
  };
};
