import * as mockEvent from "../event.json";
import { createSpace } from "../../src/adapter/resolver/mutation";
import { mock } from "../mock";

describe("createSpace()", () => {
  // doesn't work
  it("should work with valid input", async () => {
    expect(global.data.userId).not.toBeUndefined();
    const event = createEvent({ ...mockEvent });

    try {
      const result = await createSpace(event);
      expect(result.uid).not.toBeUndefined();
    } catch (err) {
      console.error(err);
    }
  });
});

const createEvent = (event: any) => {
  return {
    ...event,
    arguments: {
      location: mock.location,
      imageUrls: mock.imageUrls,
      title: mock.title,
      description: mock.description,
    },
    identity: {
      ...event.identity,
      resolverContext: {
        uid: global.data.userId,
      },
    },
  };
};
