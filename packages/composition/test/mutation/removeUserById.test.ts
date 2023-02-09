import {
  RemoveUserByIdCommand,
  RemoveUserByIdUseCase,
} from "../../src/adapter/usecase/removeUserById";

describe("removeUserById()", () => {
  it("should work with valid input", async () => {
    if (!global.data.userId) {
      throw new Error("data.userId is empty");
    }
    const input = { userId: global.data.userId };
    const usecase = new RemoveUserByIdUseCase();
    const result = await usecase.execute(new RemoveUserByIdCommand(input));
    expect(result?.uid).not.toBeUndefined();
  });
});
