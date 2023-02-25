import { faker } from "@faker-js/faker";

export const mock = {
  emailAddress: faker.internet.email(),
  firstName: faker.name.firstName(),
  lastName: faker.name.lastName(),
  password: faker.internet.password(),
  title: faker.commerce.productName(),
  imageUrls: [
    `${faker.image.imageUrl()}/${faker.random.alphaNumeric(15)}`,
    `${faker.image.imageUrl()}/${faker.random.alphaNumeric(15)}`,
  ],
  description: faker.lorem.paragraphs(),
  startDate: faker.date.past().toISOString(),
  endDate: faker.date.soon().toISOString(),
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
};
