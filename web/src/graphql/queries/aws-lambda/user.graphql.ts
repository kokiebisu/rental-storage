import { gql } from "@apollo/client";

export const PROFILE_QUERY = gql`
  query MyQuery {
    profile {
      createdAt
      emailAddress
      firstName
    }
  }
`;
