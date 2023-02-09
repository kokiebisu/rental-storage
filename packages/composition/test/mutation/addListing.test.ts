import { ListingRestClient } from "../../src/client";

describe("addListing()", () => {
  it("should work with valid input", async () => {
    expect(global.data.listingId).not.toBeUndefined();
  });

  it.todo("wrong input", () => {
    const listingClient = new ListingRestClient();
    //   title: 'HUIHIU',
    // streetAddress: 'IHI',
    // latitude: 1.5,
    // longitude: 1.5,
    // imageUrls: [ 'UIGIU' ],
    // feeType: 'MONTHLY',
    // feeAmount: 10,
    // feeCurrency: 'CAD'
    // await listingClient.addListing()
  });
});
