import gql from "graphql-tag";

export const QUERY_FIND_USER_BY_EMAIL = gql`
  query Query($email: String) {
    findUserByEmail(email: $email) {
      createdAt
      emailAddress
      firstName
      lastName
      payment {
        providerId
        providerType
      }
      uid
      updatedAt
    }
  }
`;

export const QUERY_FIND_ME = gql`
  query Query {
    findMe {
      createdAt
      emailAddress
      firstName
      lastName
    }
  }
`;

export const QUERY_FIND_USER_BY_ID = gql`
  query Query($uid: ID) {
    findUserById(uid: $uid) {
      createdAt
      emailAddress
      firstName
      lastName
      payment {
        providerId
        providerType
      }
      uid
      updatedAt
    }
  }
`;

export const MUTATION_REMOE_USER_BY_ID = gql`
  mutation MyMutation($uid: ID) {
    removeUserById(uid: $uid)
  }
`;
