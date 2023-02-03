import { CustomApolloProvider } from "./apollo";
import { AuthContextProvider } from "./auth";

export const ContextProvider = ({ children }: any) => (
  <CustomApolloProvider>
    <AuthContextProvider>{children}</AuthContextProvider>
  </CustomApolloProvider>
);
