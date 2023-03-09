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
  bookings: {
    pending: IBooking[];
    approved: IBooking[];
  };
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

interface AppsyncResolverMockEvent {
  arguments: TArguments;
  identity?: AppSyncIdentity;
  source: TSource;
  info: {
    selectionSetList: string[];
    selectionSetGraphQL: string;
    parentTypeName: string;
    fieldName: string;
    variables: { [key: string]: any };
  };
  prev: { result: { [key: string]: any } } | null;
  stash: { [key: string]: any };
}

declare const global: Global;
