import { gql } from "@apollo/client";

export default gql`
  query MyQuery {
    profile {
      createdAt
      emailAddress
      firstName
    }
  }
`;
