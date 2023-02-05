declare global {
  var data: JestMockData;
}

interface JestMockData {
  uid?: string;
  listingId?: string;
  mockFirstName: string;
  mockLastName: string;
  mockEmailAddress: string;
  mockPassword: string;
}

export {};
