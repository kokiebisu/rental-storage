import { gql } from "@apollo/client";

/**
 * AWS LAMBDA Query
 */
export const FIND_SPACES_BY_LENDER_QUERY = gql`
  query MyQuery($filter: SpaceFilter) {
    spaces(filter: $filter) {
      id
      imageUrls
      bookings {
        approved {
          createdAt
          description
        }
        pending {
          createdAt
          description
        }
      }
    }
  }
`;
