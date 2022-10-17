import gql from "graphql-tag";

export const QUERY_FIND_LISTING_BY_ID = gql`
  query Query($uid: ID!) {
    findListingById(uid: $uid) {
      latitude
      lenderId
      longitude
      streetAddress
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
      latitude
      lenderId
      longitude
      streetAddress
      uid
    }
  }
`;

export const MUTATION_ADD_LISTING = gql`
  mutation MyMutation(
    $lenderId: ID!
    $imageUrls: [String]
    $latitude: Float!
    $longitude: Float!
    $streetAddress: String!
  ) {
    addListing(
      latitude: $latitude
      lenderId: $lenderId
      longitude: $longitude
      streetAddress: $streetAddress
      imageUrls: $imageUrls
    )
  }
`;

export const MUTATION_REMOVE_LISTING_BY_ID = gql`
  mutation MyMutation($uid: ID!) {
    removeListingById(uid: $uid)
  }
`;
