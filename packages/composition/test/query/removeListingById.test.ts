import {
  RemoveListingByIdCommand,
  RemoveListingByIdUseCase,
} from "../../src/adapter/usecase/removeListingById";

describe("removeListingById()", () => {
  it("should work with valid input", async () => {
    if (!global.data.listingId) {
      throw new Error("data.listingId is empty");
    }
    const input = { listingId: global.data.listingId };
    const usecase = new RemoveListingByIdUseCase();
    const result = await usecase.execute(new RemoveListingByIdCommand(input));
    expect(result).not.toBeUndefined();
  });
});
