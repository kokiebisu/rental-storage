import { TimeUtil } from "../../src/utils";

const mock = {
  date: "Wed Oct 19 2022 23:00:12 GMT-0700 (Pacific Daylight Time)",
};

describe("LoggerUtil", () => {
  describe("toDate()", () => {
    it("should work", () => {
      const date = TimeUtil.toDate(new Date(mock.date));
      expect(date).not.toBeUndefined();
      expect(date).toBe("2022-10-20");
    });
  });
  describe("toLocaleDate()", () => {
    it("should work", () => {
      const date = TimeUtil.toDateTime(new Date(mock.date));
      expect(date).not.toBeUndefined();
      expect(date).toBe("2022-10-20 06:00:12");
    });
  });
});
