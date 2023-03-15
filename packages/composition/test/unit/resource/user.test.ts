import { UserResourceURLBuilder } from "../../../src/resource";

describe("UserResourceURLBuilder", () => {
  const builder = new UserResourceURLBuilder();

  it("finds a user by id", () => {
    const id = "123";
    const result = builder.findUser(id);
    expect(result).toEqual(`${builder.baseURL}/users/${id}`);
  });

  it("deletes a user by id", () => {
    const id = "123";
    const result = builder.deleteUser(id);
    expect(result).toEqual(`${builder.baseURL}/users/${id}`);
  });
});
