require("dotenv/config");
require("ts-node").register({
  transpileOnly: true,
});

import { faker } from "@faker-js/faker";
import { UserRestClient } from "../src/client";

let mockEmailAddress = faker.internet.email();
let mockFirstName = faker.name.firstName();
let mockLastName = faker.name.lastName();
let mockPassword = faker.internet.password();

module.exports = async function () {
  const uid = await registerUser();

  global.data = {
    uid,
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
