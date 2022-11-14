import { faker } from "@faker-js/faker";
import {
  FindListingByIdCommand,
  FindListingByIdUseCase,
} from "../../src/adapter/usecase/findListingById";

let usecase: FindListingByIdUseCase;
let mockEmailAddress = faker.internet.email();
let mockFirstName = faker.name.firstName();
let mockLastName = faker.name.lastName();
let mockPassword = faker.internet.password();

beforeAll(async () => {
  usecase = new FindListingByIdUseCase();
});

describe("findListingById()", () => {
  it("should work with valid input", async () => {
    // const commandInput = {...}
    // const command = new FindListingByIdCommand(commandInput);
    // const result = usecase.execute(command);
    // expect(result).not.toBeUndefined()
  });
});
