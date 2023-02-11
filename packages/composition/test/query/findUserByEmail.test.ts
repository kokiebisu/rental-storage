import {
  FindUserByEmailCommand,
  FindUserByEmailUseCase,
} from "../../src/adapter/usecase/findUserByEmail";

describe("findUserByEmail()", () => {
  it("should work with valid input", async () => {
    if (!global.data.mockEmailAddress) {
      throw new Error("data.mockEmailAddress is empty");
    }
    const input = { email: global.data.mockEmailAddress };
    const usecase = new FindUserByEmailUseCase();
    const result = await usecase.execute(new FindUserByEmailCommand(input));
    expect(result).not.toBeUndefined();
  });
});
