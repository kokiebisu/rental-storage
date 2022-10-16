import { AuthContextProvider } from "./auth";

export const ContextProvider = ({ children }) => (
  <AuthContextProvider>{children}</AuthContextProvider>
);
