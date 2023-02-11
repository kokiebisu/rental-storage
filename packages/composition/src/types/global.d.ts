import { BookingItem } from "../adapter/usecase/makeBooking";

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
    uid: string;
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

export {};
