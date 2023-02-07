require("dotenv/config");
require("ts-node").register({
  transpileOnly: true,
});

import { faker } from "@faker-js/faker";
import {
  BookingRestClient,
  ListingRestClient,
  UserRestClient,
} from "../src/client";

const mockEmailAddress = faker.internet.email();
const mockFirstName = faker.name.firstName();
const mockLastName = faker.name.lastName();
const mockPassword = faker.internet.password();
const mockStreetAddress = faker.address.streetAddress();
const mockLatitude = faker.address.latitude();
const mockLongitude = faker.address.longitude();
const mockImageUrls = [
  `${faker.image.imageUrl()}/${faker.random.alphaNumeric(15)}`,
  `${faker.image.imageUrl()}/${faker.random.alphaNumeric(15)}`,
];
const mockTitle = faker.company.name();
const mockFeeAmount = faker.commerce.price();
const mockFeeCurrency = faker.finance.currencyCode();
const mockFeeType = "MONTHLY";
const mockItems = [
  {
    name: faker.commerce.product(),
    imageUrls: mockImageUrls,
  },
];

module.exports = async function () {
  const userId = await registerUser();
  const listingId = await registerListing(userId);
  const bookingId = await registerBooking(userId, listingId);

  global.data = {
    userId,
    listingId,
    bookingId,
    mockEmailAddress,
    mockFirstName,
    mockLastName,
    mockPassword,
  };
};

const registerUser = async function () {
  const userClient = new UserRestClient();
  const responseData = await userClient.createUser(
    mockEmailAddress,
    mockFirstName,
    mockLastName,
    mockPassword
  );
  if (!responseData) {
    throw new Error("register user request failed");
  }
  return responseData.uid;
};

const registerListing = async function (userId: string) {
  const listingClient = new ListingRestClient();
  const responseData = await listingClient.addListing(
    userId,
    mockStreetAddress,
    Number(mockLatitude),
    Number(mockLongitude),
    mockImageUrls,
    mockTitle,
    Number(mockFeeAmount),
    mockFeeCurrency,
    mockFeeType
  );
  if (!responseData) {
    throw new Error("register listing request failed");
  }
  return responseData.uid;
};

const registerBooking = async function (userId: string, listingId: string) {
  const bookingClient = new BookingRestClient();
  const responseData = await bookingClient.makeBooking(
    userId,
    listingId,
    mockItems
  );
  if (!responseData) {
    throw new Error("register booking request failed");
  }
  return responseData.uid;
};
