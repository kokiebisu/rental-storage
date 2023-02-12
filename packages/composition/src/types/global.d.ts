import { BookingItem } from "../adapter/usecase/createBooking";
import { AppSyncResolverEvent } from "aws-lambda";

declare global {
  var data: JestMockData;

  interface Space {
    id: string;
    lenderId: string;
    streetAddress: string;
    latitude: number;
    longitude: number;
    imageUrls: string[];
    title: string;
    description: string;
  }

  interface Booking {
    id: string;
    status: string;
    userId: string;
    spaceId: string;
    items: BookingItem;
    createdAt: string;
    updatedAt: string;
  }

  interface BookingItem {
    name: string;
    imageUrls: string[];
  }

  interface User {
    id: string;
    firstName: string;
    lastName: string;
    emailAddress: string;
    streetAddress: string;
    createdAt: string;
    updatedAt: string;
  }
}

interface JestMockData {
  userId?: string;
  spaceId?: string;
  bookingId?: string;
  mockFirstName: string;
  mockLastName: string;
  mockEmailAddress: string;
  mockPassword: string;
}
