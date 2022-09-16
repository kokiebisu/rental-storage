import gql from "graphql-tag";

export const CREATE_TODO = gql`
  mutation MyMutation($description: String, $id: ID, $name: String) {
    createTodo(description: $description, id: $id, name: $name) {
      description
      id
      name
    }
  }
`;
