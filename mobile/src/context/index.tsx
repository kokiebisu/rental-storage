import { AuthContextProvider } from "./auth";
import { ProfileContextProvider } from "./profile";

export const ContextProvider = ({ children }) => (
  <ProfileContextProvider>
    <AuthContextProvider>{children}</AuthContextProvider>
  </ProfileContextProvider>
);
