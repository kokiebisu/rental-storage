import { gql } from "@apollo/client";

export const FIND_PROFILE_QUERY = gql`
  query MyQuery {
    profile {
      createdAt
      emailAddress
      firstName
    }
  }
`;
