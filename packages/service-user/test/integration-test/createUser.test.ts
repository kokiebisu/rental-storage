import { CreateUserInput, UserService } from "../../src/port";
import { UserServiceImpl } from "../../src/adapter/service";
import { faker } from "@faker-js/faker";

let service: UserService;
let mockEmailAddress = faker.internet.email();
let mockFirstName = faker.name.firstName();
let mockLastName = faker.name.lastName();
let mockPassword = faker.internet.password();

beforeAll(async () => {
  service = await UserServiceImpl.create();
});

describe("createUser()", () => {
  it("should create a user with valid input", async () => {
    const input: CreateUserInput = {
      emailAddress: mockEmailAddress,
      firstName: mockFirstName,
      lastName: mockLastName,
      password: mockPassword,
    };
    await service.createUser(input);

    const createdUser = await service.findByEmail(mockEmailAddress);
    expect(createdUser).not.toBeUndefined();
  });
});
