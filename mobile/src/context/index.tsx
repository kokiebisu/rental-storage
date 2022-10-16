import { AuthContextProvider } from "./auth";

export const ContextProvider = ({ children }) => {
  return <AuthContextProvider>{children}</AuthContextProvider>;
};
