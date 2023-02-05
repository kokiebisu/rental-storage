import { SignUpParams, useAuth } from "@/hooks/useAuth";
import { createContext } from "react";
import { User } from "../hooks/useUser";

interface AuthContext {
  user?: User | null;
  login: (user: any) => Promise<void>;
  logout: () => void;
  signup: (data: SignUpParams) => Promise<void>;
  isAuthenticated: () => boolean;
}

export const AuthContext = createContext<AuthContext>({
  user: null,
  login: async () => {},
  logout: () => {},
  signup: async () => {},
  isAuthenticated: () => false,
});

export const AuthContextProvider = ({ children }: any) => {
  const { user, signup, login, logout, isAuthenticated } = useAuth();
  return (
    <AuthContext.Provider
      value={{ user, signup, login, logout, isAuthenticated }}
    >
      {children}
    </AuthContext.Provider>
  );
};
