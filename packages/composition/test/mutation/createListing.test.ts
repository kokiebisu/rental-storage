import * as mockEvent from "../event.json";
import { createListing } from "../../src/adapter/resolver/mutation";
import { mock } from "../mock";

describe("createListing()", () => {
  it("should work with valid input", async () => {
    expect(global.data.userId).not.toBeUndefined();
    const event = createEvent({ ...mockEvent });
    const result = await createListing(event);
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
      feeAmount: mock.feeAmount,
      feeCurrency: mock.feeCurrency,
      feeType: mock.feeType,
    },
    identity: {
      ...event.identity,
      resolverContext: {
        uid: global.data.userId,
      },
    },
  };
};
