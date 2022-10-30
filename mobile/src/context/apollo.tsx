import {
  ApolloClient,
  ApolloLink,
  ApolloProvider,
  HttpLink,
  InMemoryCache,
  NormalizedCacheObject,
} from "@apollo/client";
import { AuthOptions, createAuthLink } from "aws-appsync-auth-link";
import { createSubscriptionHandshakeLink } from "aws-appsync-subscription-link";

import { appsyncConfig } from "../env";
import { useContext } from "react";
import { AuthContext } from "./auth";

export const CustomApolloProvider = ({ children }: any) => {
  const url = appsyncConfig.GRAPHQL_ENDPOINT;
  const region = appsyncConfig.REGION;
  const { authState } = useContext(AuthContext);

  const auth = {
    type: appsyncConfig.AUTHENTICATION_TYPE,
    token: async () => {
      try {
        const token = authState.userToken;
        return token;
      } catch (err) {
        console.error(err);
      }
    },
  } as AuthOptions;

  const httpLink = new HttpLink({ uri: url });

  const link = ApolloLink.from([
    createAuthLink({ url, region: appsyncConfig.REGION, auth }),
    createSubscriptionHandshakeLink({ url, region, auth }, httpLink),
  ]);

  const client: ApolloClient<NormalizedCacheObject> = new ApolloClient({
    link,
    cache: new InMemoryCache(),
    connectToDevTools: true,
  });

  return <ApolloProvider client={client}>{children}</ApolloProvider>;
};
