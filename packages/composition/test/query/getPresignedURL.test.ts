import { faker } from "@faker-js/faker";
import {
  AddListingCommand,
  AddListingUseCase,
} from "../../src/adapter/usecase/addListing";
import {
  GetPresignedURLCommand,
  GetPresignedURLUseCase,
} from "../../src/adapter/usecase/getPresignedURL";

let usecase: GetPresignedURLUseCase;
let mockEmailAddress = faker.internet.email();
let mockFirstName = faker.name.firstName();
let mockLastName = faker.name.lastName();
let mockPassword = faker.internet.password();
let commandInput: GetPresignedURLCommand;

beforeAll(async () => {
  usecase = new GetPresignedURLUseCase();
});

describe("getPresignedURL()", () => {
  it("should work with valid input", async () => {
    // const commandInput = {...}
    // const command = new GetPresignedURLCommand(commandInput);
    // const result = usecase.execute(command);
    // expect(result).not.toBeUndefined()
  });
});
