import gql from "graphql-tag";

export const GET_TODOS = gql`
  query MyQuery {
    getTodos {
      id
      description
      name
    }
  }
`;
