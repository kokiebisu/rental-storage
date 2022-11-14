import { faker } from "@faker-js/faker";
import {
  FindMyListingsCommand,
  FindMyListingsUseCase,
} from "../../src/adapter/usecase/findMyListings";

let usecase: FindMyListingsUseCase;
let mockEmailAddress = faker.internet.email();
let mockFirstName = faker.name.firstName();
let mockLastName = faker.name.lastName();
let mockPassword = faker.internet.password();

beforeAll(async () => {
  usecase = new FindMyListingsUseCase();
});

describe("findMyListings()", () => {
  it("should work with valid input", async () => {
    // const commandInput = {...}
    // const command = new FindMyListingsCommand(commandInput)
    // const result = usecase.execute(command)
    // expect(result).not.toBeUndefined()
  });
});
