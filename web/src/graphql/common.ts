import { gql } from "@apollo/client";

gql`
  input CoordinateInput {
    latitude: Float
    longitude: Float
  }

  input LocationInput {
    address: String
    city: String
    country: String
    countryCode: String
    phone: String
    province: String
    provinceCode: String
    zip: String
    coordinate: CoordinateInput
  }
`;
