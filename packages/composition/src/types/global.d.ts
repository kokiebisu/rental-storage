import { BookingItem } from "../adapter/usecase/createBooking";
import { AppSyncResolverEvent } from "aws-lambda";

declare global {
  var data: JestMockData;

  interface Listing {
    uid: string;
    lenderId: string;
    streetAddress: string;
    latitude: number;
    longitude: number;
    imageUrls: string[];
    feeAmount: number[];
    feeCurrency: number[];
    feeType: number[];
  }

  interface Booking {
    id: string;
    status: string;
    userId: string;
    listingId: string;
    items: BookingItem;
    createdAt: string;
    updatedAt: string;
  }

  interface BookingItem {
    name: string;
    imageUrls: string[];
  }

  interface User {
    uid: string;
    firstName: string;
    lastName: string;
    emailAddress: string;
    streetAddress: string;
  }
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
