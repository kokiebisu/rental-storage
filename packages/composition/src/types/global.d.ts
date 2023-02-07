declare global {
  var data: JestMockData;
}

interface JestMockData {
  userId?: string;
  listingId?: string;
  bookingId?: string;
  mockFirstName: string;
  mockLastName: string;
  mockEmailAddress: string;
  mockPassword: string;
}

export {};
