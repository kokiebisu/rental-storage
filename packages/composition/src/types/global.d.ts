declare global {
  const data: JestMockData;
}

interface JestMockData {
  uid?: string;
  mockFirstName: string;
  mockLastName: string;
  mockEmailAddress: string;
  mockPassword: string;
}

export {};
