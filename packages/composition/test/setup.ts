require("dotenv/config");
require("ts-node").register({
  transpileOnly: true,
});

import { faker } from "@faker-js/faker";
import {
  BookingRestClient,
  SpaceRestClient,
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
const mockDescription = faker.lorem.paragraph();
const mockTitle = faker.company.name();
const mockStartDate = faker.date.past();
const mockEndDate = faker.date.soon();

module.exports = async function () {
  const userId = await registerUser();
  const spaceId = await registerSpace(userId);
  const bookingId = await registerBooking(userId, spaceId);

  global.data = {
    userId,
    spaceId,
    bookingId,
    mockEmailAddress,
    mockFirstName,
    mockLastName,
    mockPassword,
    mockImageUrls,
    mockDescription,
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

const registerSpace = async function (userId: string) {
  const spaceClient = new SpaceRestClient();
  try {
    const responseData = await spaceClient.createSpace(
      userId,
      mockStreetAddress,
      Number(mockLatitude),
      Number(mockLongitude),
      mockImageUrls,
      mockTitle,
      mockDescription
    );
    if (!responseData) {
      throw new Error(
        `register space request failed with latitude: ${mockLatitude}, longitude: ${mockLongitude}`
      );
    }
    return responseData.uid;
  } catch (err) {
    console.error(err);
  }
};

const registerBooking = async function (userId: string, spaceId: string) {
  const bookingClient = new BookingRestClient();
  const responseData = await bookingClient.createBooking(
    userId,
    spaceId,
    mockImageUrls,
    mockDescription,
    mockStartDate.toISOString(),
    mockEndDate.toISOString()
  );
  if (!responseData) {
    throw new Error("register booking request failed");
  }
  return responseData.uid;
};
