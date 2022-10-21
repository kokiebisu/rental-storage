import gql from "graphql-tag";

export const QUERY_FIND_LISTING_BY_ID = gql`
  query Query($uid: ID!) {
    findListingById(uid: $uid) {
      fee {
        amount
        currency
      }
      imageUrls
      latitude
      lenderId
      longitude
      streetAddress
      title
      uid
    }
  }
`;

export const QUERY_FIND_MY_LISTINGS = gql`
  query Query {
    findMyListings {
      fee {
        amount
        currency
      }
      imageUrls
      latitude
      lenderId
      longitude
      streetAddress
      title
      uid
    }
  }
`;

export const QUERY_FIND_LISTINGS_WITHIN_LAT_LNG = gql`
  query Query($latitude: Float!, $longitude: Float!, $range: Int!) {
    findListingsWithinLatLng(
      latitude: $latitude
      longitude: $longitude
      range: $range
    ) {
      fee {
        amount
        currency
      }
      imageUrls
      latitude
      lenderId
      longitude
      streetAddress
      title
      uid
    }
  }
`;

export const MUTATION_ADD_LISTING = gql`
  mutation Mutation(
    $feeAmount: Int!
    $feeCurrency: String!
    $imageUrls: [String]
    $latitude: Float!
    $longitude: Float!
    $streetAddress: String!
    $title: String!
  ) {
    addListing(
      fee: { amount: $feeAmount, currency: $feeCurrency }
      latitude: $latitude
      longitude: $longitude
      streetAddress: $streetAddress
      title: $title
      imageUrls: $imageUrls
    )
  }
`;

export const MUTATION_REMOVE_LISTING_BY_ID = gql`
  mutation Mutation($uid: ID!) {
    removeListingById(uid: $uid)
  }
`;
