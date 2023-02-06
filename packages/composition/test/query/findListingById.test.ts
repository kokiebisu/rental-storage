import {
  FindListingByIdCommand,
  FindListingByIdUseCase,
} from "../../src/adapter/usecase/findListingById";

describe("findListingById()", () => {
  it("should work with valid input", async () => {
    if (!global.data.listingId) {
      throw new Error("data.listingId is empty");
    }
    const input = { listingId: global.data.listingId };
    const usecase = new FindListingByIdUseCase();
    const result = await usecase.execute(new FindListingByIdCommand(input));
    expect(result.listing).not.toBeUndefined();
  });
});
