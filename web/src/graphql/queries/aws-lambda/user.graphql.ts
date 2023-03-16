import { gql } from "@apollo/client";

export const PROFILE_QUERY = gql`
  query MyQuery {
    profile {
      id
    }
  }
`;

export const FIND_MY_BOOKINGS_QUERY = gql`
  query MyQuery {
    profile {
      bookings {
        pending {
          id
          imageUrls
          description
          startDate
          endDate
        }
        approved {
          id
          imageUrls
          description
          startDate
          endDate
        }
      }
    }
  }
`;
