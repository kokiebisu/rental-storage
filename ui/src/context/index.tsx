import { TodoContextProvider } from "./todos";

export const ContextProvider = ({ children }) => {
  return <TodoContextProvider>{children}</TodoContextProvider>;
};
