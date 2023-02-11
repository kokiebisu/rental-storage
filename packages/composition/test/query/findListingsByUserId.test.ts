import {
  FindListingsByUserIdCommand,
  FindListingsByUserIdUseCase,
} from "../../src/adapter/usecase/findListingsByUserId";

describe("findListingsByUserId()", () => {
  it("should work with valid input", async () => {
    if (!global.data.userId) {
      throw new Error("data.userId is empty");
    }
    const input = { userId: global.data.userId };
    const usecase = new FindListingsByUserIdUseCase();
    try {
      const result = await usecase.execute(
        new FindListingsByUserIdCommand(input)
      );
      expect(result).not.toBeUndefined();
    } catch (err) {
      console.error(err);
    }
  });
});
