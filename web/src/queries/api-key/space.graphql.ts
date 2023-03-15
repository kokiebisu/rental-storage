import { gql } from "@apollo/client";

/**
 * API KEY Query
 */
export const FIND_SPACES_QUERY = gql`
  query MyQuery($filter: SpaceFilter) {
    spaces(filter: $filter) {
      id
      title
      imageUrls
      location {
        address
      }
    }
  }
`;

/**
 * API KEY
 */
export const FIND_SPACE_QUERY = gql`
  query MyQuery($id: ID) {
    space(id: $id) {
      createdAt
      description
      id
      lenderId
      imageUrls
      title
      updatedAt
    }
  }
`;
