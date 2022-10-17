import { AuthContextProvider } from "./auth";
import { ProfileContextProvider } from "./profile";

export const ContextProvider = ({ children }) => (
  <AuthContextProvider>
    <ProfileContextProvider>{children}</ProfileContextProvider>
  </AuthContextProvider>
);
