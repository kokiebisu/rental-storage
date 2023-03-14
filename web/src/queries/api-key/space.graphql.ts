import { gql } from "@apollo/client";

/**
 * API KEY Query
 */
export const FIND_SPACES_QUERY = gql`
  query MyQuery($userId: ID) {
    spaces(userId: $userId) {
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
