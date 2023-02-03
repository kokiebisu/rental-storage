import { CustomApolloProvider } from "./apollo";

export const ContextProvider = ({ children }: any) => (
    <CustomApolloProvider>
        {children}
    </CustomApolloProvider>
);
