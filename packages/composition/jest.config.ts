module.exports = {
  roots: ["<rootDir>/test"],
  transform: {
    "^.+\\.ts?$": "ts-jest",
  },
  testRegex: "(/__tests__/.*|(\\.|/)(test|spec))\\.ts?$",
  moduleFileExtensions: ["ts", "js", "json", "node"],
  collectCoverage: true,
  clearMocks: true,
  coverageDirectory: "coverage",
  setupFiles: ["dotenv/config"],
  testTimeout: 100000,
  globalSetup: "./test/setup.ts",
};
