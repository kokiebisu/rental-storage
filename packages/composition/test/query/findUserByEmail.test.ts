import { faker } from "@faker-js/faker";
import {
  AddListingCommand,
  AddListingUseCase,
} from "../../src/adapter/usecase/addListing";
import {
  FindUserByEmailCommand,
  FindUserByEmailUseCase,
} from "../../src/adapter/usecase/findUserByEmail";

let usecase: FindUserByEmailUseCase;
let mockEmailAddress = faker.internet.email();
let mockFirstName = faker.name.firstName();
let mockLastName = faker.name.lastName();
let mockPassword = faker.internet.password();

beforeAll(async () => {
  usecase = new FindUserByEmailUseCase();
});

describe("findUserByEmail()", () => {
  it("should work with valid input", async () => {
    // const commandInput = {...}
    // const command = new FindUserByEmailCommand(commandInput);
    // const result = usecase.execute(command);
    // expect(result).not.toBeUndefined()
  });
});
