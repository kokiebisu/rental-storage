import { SpaceResourceURLBuilder } from "../../../src/resource";

describe("SpaceResourceURLBuilder", () => {
  const builder = new SpaceResourceURLBuilder();
  it("generates correct url to find a space by id", () => {
    const id = "123";
    const result = builder.findSpace(id);
    expect(result).toEqual(`${builder.baseURL}/spaces/${id}`);
  });

  it("generates correct url to find spaces by userId", () => {
    const userId = "123";
    const result = builder.findSpaces({ userId });
    expect(result).toEqual(`${builder.baseURL}/spaces?userId=${userId}`);
  });

  it("generates correct url to find spaces by limit and offset", () => {
    const offset = 0;
    const limit = 10;
    const result = builder.findSpaces({ limit, offset });
    expect(result).toEqual(
      `${builder.baseURL}/spaces?offset=${offset}&limit=${limit}`
    );
  });
});
