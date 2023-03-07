import { ImageResourceURLBuilder } from "../../../src/client";

describe("ImageResourceURLBuilder", () => {
  const builder = new ImageResourceURLBuilder();
  it("creates a presigned url", () => {
    const filename = "123";
    const result = builder.getPresignedURL(filename);
    expect(result).toEqual(`${builder.baseURL}/images/${filename}`);
  });
});
