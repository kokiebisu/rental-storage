import { faker } from "@faker-js/faker";
import { FindListingsWithinLatLngUseCase } from "../../src/adapter/usecase/findListingsWithinLatLng";

let usecase: FindListingsWithinLatLngUseCase;
let mockEmailAddress = faker.internet.email();
let mockFirstName = faker.name.firstName();
let mockLastName = faker.name.lastName();
let mockPassword = faker.internet.password();

beforeAll(async () => {
  usecase = new FindListingsWithinLatLngUseCase();
});

describe("findListingsWithinLatLng()", () => {
  it("should work with valid input", async () => {
    // const commandInput = {...}
    // const command = new FindListingsWithinLatLngUseCase(commandInput);
    // const result = usecase.execute(command);
    // expect(result).not.toBeUndefined()
  });
});
