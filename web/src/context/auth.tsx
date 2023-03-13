import { createContext } from "react";

import { SignUpParams, useAuth } from "@/hooks/useAuth";
import { User } from "@/hooks/useUser";

interface AuthContext {
  user?: User | null;
  login: (user: any) => Promise<void>;
  logout: () => void;
  signup: (data: SignUpParams) => Promise<void>;
  checkIsAuthenticated: () => boolean;
}

export const AuthContext = createContext<AuthContext>({
  user: null,
  login: async () => {},
  logout: () => {},
  signup: async () => {},
  checkIsAuthenticated: () => false,
});

export const AuthContextProvider = ({ children }: any) => {
  const { user, signup, login, logout, checkIsAuthenticated } = useAuth();
  return (
    <AuthContext.Provider
      value={{ user, signup, login, logout, checkIsAuthenticated }}
    >
      {children}
    </AuthContext.Provider>
  );
};
