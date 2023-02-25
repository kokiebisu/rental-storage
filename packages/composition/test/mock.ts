import { faker } from "@faker-js/faker";

export const mock = {
  emailAddress: faker.internet.email(),
  firstName: faker.name.firstName(),
  lastName: faker.name.lastName(),
  password: faker.internet.password(),
  title: faker.company.name(),
  imageUrls: `[
    "${faker.image.imageUrl()}/${faker.random.alphaNumeric(15)}",
    "${faker.image.imageUrl()}/${faker.random.alphaNumeric(15)}",
  ]`,
  description: faker.lorem.paragraphs(),
  startDate: faker.date.past(),
  endDate: faker.date.soon(),
  location: {
    address: faker.address.streetAddress(),
    city: faker.address.cityName(),
    country: faker.address.country(),
    countryCode: faker.address.countryCode(),
    phone: faker.phone.number("##########"),
    province: "CA",
    zip: faker.address.zipCode(),
    coordinate: {
      latitude: Number(faker.address.latitude()),
      longitude: Number(faker.address.longitude()),
    },
  },
};
