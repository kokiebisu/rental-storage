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

const mock = {
  emailAddress: faker.internet.email(),
  firstName: faker.name.firstName(),
  lastName: faker.name.lastName(),
  password: faker.internet.password(),
  location: {
    address: faker.address.streetAddress(),
    city: faker.address.cityName(),
    country: faker.address.country(),
    countryCode: faker.address.countryCode(),
    phone: faker.phone.number("##########"),
    province: "British Columbia",
    provinceCode: "BC",
    zip: faker.address.zipCode(),
    coordinate: {
      latitude: Number(faker.address.latitude()),
      longitude: Number(faker.address.longitude()),
    },
  },
  imageUrls: [
    `${faker.image.imageUrl()}/${faker.random.alphaNumeric(15)}`,
    `${faker.image.imageUrl()}/${faker.random.alphaNumeric(15)}`,
  ],
  description: faker.lorem.paragraph(),
  title: faker.commerce.productName(),
  startDate: faker.date.past().toISOString(),
  endDate: faker.date.soon().toISOString(),
};

module.exports = async function () {
  const userId = await registerUser();
  const spaceId = await registerSpace(userId);
  const bookingId = await registerBooking(userId, spaceId);

  global.data = {
    userId,
    spaceId,
    bookingId,
    emailAddress: mock.emailAddress,
    firstName: mock.firstName,
    lastName: mock.lastName,
    password: mock.password,
    imageUrls: mock.imageUrls,
    description: mock.description,
  };
};

const registerUser = async function () {
  const userClient = new UserRestClient();
  const responseData = await userClient.createUser(
    mock.emailAddress,
    mock.firstName,
    mock.lastName,
    mock.password
  );
  if (!responseData) {
    throw new Error("register user request failed");
  }
  return responseData.uid;
};

const registerSpace = async function (userId: string) {
  const spaceClient = new SpaceRestClient();

  const responseData = await spaceClient.createSpace(
    userId,
    mock.imageUrls,
    mock.title,
    mock.description,
    mock.location
  );
  if (!responseData) {
    throw new Error(
      `register space request failed with latitude: ${mock.location.coordinate.latitude}, longitude: ${mock.location.coordinate.longitude}`
    );
  }
  return responseData.uid;
};

const registerBooking = async function (userId: string, spaceId: string) {
  const bookingClient = new BookingRestClient();
  const responseData = await bookingClient.createBooking(
    userId,
    spaceId,
    mock.imageUrls,
    mock.description,
    mock.startDate,
    mock.endDate
  );
  if (!responseData) {
    throw new Error("register booking request failed");
  }
  return responseData.uid;
};
