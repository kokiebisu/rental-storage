require("dotenv/config");
require("ts-node").register({
  transpileOnly: true,
});

import { faker } from "@faker-js/faker";
import { ListingRestClient, UserRestClient } from "../src/client";

const mockEmailAddress = faker.internet.email();
const mockFirstName = faker.name.firstName();
const mockLastName = faker.name.lastName();
const mockPassword = faker.internet.password();
const mockStreetAddress = faker.address.streetAddress();
const mockLatitude = faker.address.latitude();
const mockLongitude = faker.address.longitude();
const mockImageUrls = [
  `${faker.image.imageUrl()}/${faker.random.alphaNumeric(20)}`,
  `${faker.image.imageUrl()}/${faker.random.alphaNumeric(20)}`,
];
const mockTitle = faker.company.name();
const mockFeeAmount = faker.commerce.price();
const mockFeeCurrency = faker.finance.currencyCode();
const mockFeeType = "MONTHLY";

module.exports = async function () {
  const uid = await registerUser();
  const listingId = await registerListing(uid);

  global.data = {
    uid,
    listingId,
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
  return responseData.uid;
};
