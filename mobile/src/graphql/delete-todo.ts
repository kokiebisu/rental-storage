import gql from "graphql-tag";

export const DELETE_TODO = gql`
  mutation MyMutation($description: String, $id: ID, $name: String, $id1: ID) {
    deleteTodo(id: $id1) {
      id
    }
  }
`;
