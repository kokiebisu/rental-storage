import { faker } from "@faker-js/faker";
import {
  AddListingCommand,
  AddListingUseCase,
} from "../../src/adapter/usecase/addListing";

let usecase: AddListingUseCase;
let mockEmailAddress = faker.internet.email();
let mockFirstName = faker.name.firstName();
let mockLastName = faker.name.lastName();
let mockPassword = faker.internet.password();

beforeAll(async () => {
  usecase = new AddListingUseCase();
});

describe("makeBooking()", () => {
  it("should work with valid input", async () => {
    // const commandInput = {...}
    // const command = new AddListingCommand(commandInput)
    // const result = usecase.execute(command)
    // expect(result).not.toBeUndefined()
  });
});
