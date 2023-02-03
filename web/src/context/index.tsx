import { CustomApolloProvider } from "./apollo";

export const ContextProvider = ({ children }) => (
    <CustomApolloProvider>
        {children}
    </CustomApolloProvider>
);
