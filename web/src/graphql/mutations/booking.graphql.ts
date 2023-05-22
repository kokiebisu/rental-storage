import { gql } from "@apollo/client";

export const MAKE_BOOKING_MUTATION = gql`
  mutation Mutation(
    $description: String = ""
    $endDate: String = ""
    $imageUrls: [String!] = ""
    $spaceId: ID = ""
    $startDate: String = ""
  ) {
    bookingCreate(
      description: $description
      endDate: $endDate
      imageUrls: $imageUrls
      spaceId: $spaceId
      startDate: $startDate
    ) {
      uid
    }
  }
`;
