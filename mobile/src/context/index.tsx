import { ApolloContextProvider } from "./apollo";
import { AuthContextProvider } from "./auth";

export const ContextProvider = ({ children }) => (
  <AuthContextProvider>
    <ApolloContextProvider>{children}</ApolloContextProvider>
  </AuthContextProvider>
);
