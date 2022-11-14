import { faker } from "@faker-js/faker";
import {
  FindAllCreatedBookingsCommand,
  FindAllCreatedBookingsUseCase,
} from "../../src/adapter/usecase/findAllCreatedBookings";

let usecase: FindAllCreatedBookingsUseCase;
let mockEmailAddress = faker.internet.email();
let mockFirstName = faker.name.firstName();
let mockLastName = faker.name.lastName();
let mockPassword = faker.internet.password();

beforeAll(async () => {
  usecase = new FindAllCreatedBookingsUseCase();
});

describe("findAllCreatedBookings()", () => {
  it("should work with valid input", async () => {
    // const commandInput = {...}
    // const command = new FindAllCreatedBookingsCommand(commandInput)
    // const result = usecase.execute(command)
    // expect(result).not.toBeUndefined()
  });
});
