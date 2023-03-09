require("dotenv/config");
require("ts-node").register({
  transpileOnly: true,
});

import { faker } from "@faker-js/faker";
import { CreateBookingCommand } from "../../src/adapter/usecase/createBooking";
import { CreateSpaceCommand } from "../../src/adapter/usecase/createSpace";
import {
  BookingResourceURLBuilder,
  RestAPIClient,
  SpaceResourceURLBuilder,
  UserResourceURLBuilder,
} from "../../src/client";

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
  const client = new RestAPIClient();
  const builder = new UserResourceURLBuilder();

  const response = await client.post<
    { uid: string },
    {
      emailAddress: string;
      firstName: string;
      lastName: string;
      password: string;
    }
  >(builder.createUser(), {
    emailAddress: mock.emailAddress,
    firstName: mock.firstName,
    lastName: mock.lastName,
    password: mock.password,
  });
  if (!response.data) {
    throw new Error("register user request failed");
  }
  return response.data.uid;
};

const registerSpace = async function (userId: string) {
  const client = new RestAPIClient();
  const builder = new SpaceResourceURLBuilder();

  const response = await client.post<{ uid: string }, CreateSpaceCommand>(
    builder.createSpace(),
    {
      lenderId: userId,
      location: mock.location,
      imageUrls: mock.imageUrls,
      title: mock.title,
      description: mock.description,
    }
  );
  if (!response.data) {
    throw new Error(
      `register space request failed with latitude: ${mock.location.coordinate.latitude}, longitude: ${mock.location.coordinate.longitude}`
    );
  }
  return response.data.uid;
};

const registerBooking = async function (userId: string, spaceId: string) {
  const client = new RestAPIClient();
  const builder = new BookingResourceURLBuilder();
  const response = await client.post<{ uid: string }, CreateBookingCommand>(
    builder.createBooking(),
    {
      userId,
      spaceId,
      imageUrls: mock.imageUrls,
      description: mock.description,
      startDate: mock.startDate,
      endDate: mock.endDate,
    }
  );
  if (!response.data) {
    throw new Error("register booking request failed");
  }
  return response.data.uid;
};
