import { faker } from "@faker-js/faker";
import { FindMeUseCase } from "../../src/adapter/usecase/findMe";

let usecase: FindMeUseCase;
let mockEmailAddress = faker.internet.email();
let mockFirstName = faker.name.firstName();
let mockLastName = faker.name.lastName();
let mockPassword = faker.internet.password();

beforeAll(async () => {
  usecase = new FindMeUseCase();
});

describe("findMe()", () => {
  it("should work with valid input", async () => {
    // const commandInput = {...}
    // const command = new FindMeUseCase(commandInput)
    // const result = usecase.execute(command)
    // expect(result).not.toBeUndefined()
  });
});
