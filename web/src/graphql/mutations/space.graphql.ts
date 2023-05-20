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

export const SPACE_CREATE_MUTATION = gql`
  mutation MyMutation(
    $title: String
    $location: LocationInput
    $imageUrls: [String]
    $description: String
  ) {
    spaceCreate(
      title: $title
      location: $location
      imageUrls: $imageUrls
      description: $description
    ) {
      uid
    }
  }
`;
