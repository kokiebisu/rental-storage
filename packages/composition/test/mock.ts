import { faker } from "@faker-js/faker";

export const mock = {
  emailAddress: faker.internet.email(),
  firstName: faker.name.firstName(),
  lastName: faker.name.lastName(),
  password: faker.internet.password(),
  streetAddress: faker.address.streetAddress(),
  latitude: Number(faker.address.latitude()),
  longitude: Number(faker.address.longitude()),
  title: faker.company.name(),
  imageUrls: `[
    "${faker.image.imageUrl()}/${faker.random.alphaNumeric(15)}",
    "${faker.image.imageUrl()}/${faker.random.alphaNumeric(15)}",
  ]`,
  description: faker.lorem.paragraphs(),
  startDate: faker.date.past(),
  endDate: faker.date.soon(),
};
