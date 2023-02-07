import { FindMeCommand, FindMeUseCase } from "../../src/adapter/usecase/findMe";

describe("findMe()", () => {
  it("should work with valid input", async () => {
    if (!global.data.userId) {
      throw new Error("data.uid is empty");
    }
    const input = { userId: global.data.userId };
    const usecase = new FindMeUseCase();
    const result = await usecase.execute(new FindMeCommand(input));
    expect(result.user).not.toBeUndefined();
  });
});
