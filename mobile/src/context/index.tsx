import { CustomApolloProvider } from "./apollo";
import { AuthContextProvider } from "./auth";
import { ProfileContextProvider } from "./profile";

export const ContextProvider = ({ children }) => (
  <AuthContextProvider>
    <CustomApolloProvider>
      <ProfileContextProvider>{children}</ProfileContextProvider>
    </CustomApolloProvider>
  </AuthContextProvider>
);
