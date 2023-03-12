import { SpaceResourceURLBuilder } from "../../../src/resource";

describe("SpaceResourceURLBuilder", () => {
  const builder = new SpaceResourceURLBuilder();
  it("finds a space by id", () => {
    const id = "123";
    const result = builder.findSpace(id);
    expect(result).toEqual(`${builder.baseURL}/spaces/${id}`);
  });

  it("finds spaces by userId", () => {
    const userId = "123";
    const result = builder.findSpaces({ userId });
    expect(result).toEqual(`${builder.baseURL}/spaces?userId=${userId}`);
  });
});
