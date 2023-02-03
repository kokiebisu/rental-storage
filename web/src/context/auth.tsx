export const AuthContext = createContext({
  authState: {
    isLoading: false,
  },
  signOut: () => Promise<void>,
  signIn: (data: any) => Promise<void>,
  signUp: (data: any) => Promise<void>,
});
