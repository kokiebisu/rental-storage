declare interface Global {
  data: JestMockData;
}

interface JestMockData {
  userId?: string;
  spaceId?: string;
  bookingId?: string;
  firstName: string;
  lastName: string;
  emailAddress: string;
  password: string;
  imageUrls: string[];
  description: string;
}

interface ISpace {
  id: string;
  lenderId: string;
  location: ILocation;
  imageUrls: string[];
  title: string;
  description: string;
  createdAt: string;
  updatedAt: string;
}

interface IBooking {
  id: string;
  status: string;
  imageUrls: string[];
  description: string;
  userId: string;
  spaceId: string;
  createdAt: string;
  updatedAt: string;
  startDate: string;
  endDate: string;
}

interface IUser {
  id: string;
  firstName: string;
  lastName: string;
  emailAddress: string;
  streetAddress: string;
  createdAt: string;
  updatedAt: string;
}

interface ILocation {
  address: string;
  city: string;
  country: string;
  countryCode: string;
  phone: string;
  province: string;
  provinceCode: string;
  zip: string;
  coordinate: Coordinate;
}

interface ICoordinate {
  latitude: number;
  longitude: number;
}

declare const global: Global;
