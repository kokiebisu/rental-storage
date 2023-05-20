import { gql } from "@apollo/client";

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
