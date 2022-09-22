import { useMutation } from "@apollo/client";
import { createContext, useState } from "react";
import { GET_TODOS } from "../graphql/all-todo";
import { CREATE_TODO } from "../graphql/create-todo";
import { DELETE_TODO } from "../graphql/delete-todo";

const TodoContext = createContext(null);

export const TodoContextProvider = ({ children }) => {
  const [addTodoMutationFunction, { error: createError }] =
    useMutation(CREATE_TODO);

  const [deleteTodoMutateFunction] = useMutation(DELETE_TODO, {
    refetchQueries: [GET_TODOS, "listTodos"],
  });

  const createTodo = async (todo) => {
    try {
      addTodoMutationFunction({ variables: { input: { todo } } });
    } catch (err) {
      console.error("error creating todo: ", err);
    }
  };

  const deleteTodo = async (id) => {
    try {
      deleteTodoMutateFunction({ variables: { input: { id } } });
    } catch (err) {
      console.error("error deleting todo: ", err);
    }
  };

  return (
    <TodoContext.Provider value={{ createTodo, deleteTodo }}>
      {children}
    </TodoContext.Provider>
  );
};
