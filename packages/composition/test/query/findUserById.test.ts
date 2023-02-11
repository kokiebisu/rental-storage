import {
  FindUserByIdCommand,
  FindUserByIdUseCase,
} from "../../src/adapter/usecase/findUserById";

describe("findUserById()", () => {
  it("should work with valid input", async () => {
    if (!global.data.userId) {
      throw new Error("data.userId is empty");
    }
    const input = { userId: global.data.userId };
    const usecase = new FindUserByIdUseCase();
    const result = await usecase.execute(new FindUserByIdCommand(input));
    expect(result).not.toBeUndefined();
  });
});
