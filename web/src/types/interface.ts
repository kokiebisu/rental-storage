export interface Space {
  id: string;
  title: string;
  imageUrls: string[];
  location: Location;
}

interface Location {
  address: string;
  coordinate: Coordinate;
}

interface Coordinate {
  latitude: number;
  longitude: number;
}

export interface SignUpParams {
  firstName: string;
  lastName: string;
  emailAddress: string;
  password: string;
}

export interface SignInParams {
  emailAddress: string;
  password: string;
}

export interface User {
  id: string;
  imageUrl?: string;
}

export interface Booking {
  id: string;
  imageUrls: string;
  description: string;
  startDate: string;
  endDate: string;
}
