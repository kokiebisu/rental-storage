import { AuthContextProvider } from "./auth";

export const ContextProvider = ({ children }: any) => (
  <AuthContextProvider>{children}</AuthContextProvider>
);
