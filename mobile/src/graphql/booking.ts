import gql from "graphql-tag";

export const MUTATION_MAKE_BOOKING = gql`
  mutation MyMutation(
    $amount: Int
    $currency: String
    $imageUrls1: [String]
    $name: String
    $listingId: ID
    $userId: ID
  ) {
    makeBooking(
      amount: $amount
      currency: $currency
      listingId: $listingId
      userId: $userId
      items: { imageUrls: $imageUrls1, name: $name }
    )
  }
`;
