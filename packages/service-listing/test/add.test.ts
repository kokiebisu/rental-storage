const add = () => {
  return 1 + 3;
};

describe("Test", () => {
  it("should work", () => {
    expect(add()).toEqual(4);
  });
});
