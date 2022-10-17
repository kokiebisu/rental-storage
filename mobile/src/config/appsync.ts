import { createContext, useContext } from "react";
import {
  ApolloClient,
  ApolloLink,
  HttpLink,
  InMemoryCache,
  NormalizedCacheObject,
} from "@apollo/client";
import { AuthOptions, createAuthLink } from "aws-appsync-auth-link";
import { createSubscriptionHandshakeLink } from "aws-appsync-subscription-link";
import { appsyncConfig } from "../env";
import * as SecureStore from "expo-secure-store";

const url = appsyncConfig.GRAPHQL_ENDPOINT;
const region = appsyncConfig.REGION;
const auth = {
  type: appsyncConfig.AUTHENTICATION_TYPE,
  token: async () => {
    try {
      return await SecureStore.getItemAsync("userToken");
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

const Client: ApolloClient<NormalizedCacheObject> = new ApolloClient({
  link,
  cache: new InMemoryCache(),
  connectToDevTools: true,
});

export { Client };
