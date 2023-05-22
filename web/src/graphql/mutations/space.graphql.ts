import { gql } from "@apollo/client";

export const SPACE_CREATE_MUTATION = gql`
  mutation SpaceCreateMutation(
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
