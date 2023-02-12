import * as mockEvent from "../event.json";
import { createSpace } from "../../src/adapter/resolver/mutation";
import { mock } from "../mock";

describe("createSpace()", () => {
  it("should work with valid input", async () => {
    expect(global.data.userId).not.toBeUndefined();
    const event = createEvent({ ...mockEvent });
    const result = await createSpace(event);
    expect(result.uid).not.toBeUndefined();
  });
});

const createEvent = (event: any) => {
  return {
    ...event,
    arguments: {
      streetAddress: mock.streetAddress,
      latitude: mock.latitude,
      longitude: mock.longitude,
      imageUrls: mock.imageUrls,
      title: mock.title,
    },
    identity: {
      ...event.identity,
      resolverContext: {
        uid: global.data.userId,
      },
    },
  };
};
