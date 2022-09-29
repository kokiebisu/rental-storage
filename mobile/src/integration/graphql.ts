import { AuthOptions, createAuthLink } from "aws-appsync-auth-link";
import { createSubscriptionHandshakeLink } from "aws-appsync-subscription-link";

import {
  ApolloProvider,
  ApolloClient,
  InMemoryCache,
  HttpLink,
  ApolloLink,
} from "@apollo/client";

// import AWSAppSyncClient from "aws-appsync";
// import { appsyncConfig } from "../env/appsync";

// export const appsyncClient = new AWSAppSyncClient({
//   url: appsyncConfig.GRAPHQL_ENDPOINT,
//   region: appsyncConfig.REGION,
//   auth: {
//     type: appsyncConfig.AUTHENTICATION_TYPE as any,
//     apiKey: appsyncConfig.API_KEY,
//   },
// });

import { appsyncConfig } from "../env";

const url = appsyncConfig.GRAPHQL_ENDPOINT;
const region = appsyncConfig.REGION;
const auth = {
  type: appsyncConfig.AUTHENTICATION_TYPE,
  apiKey: appsyncConfig.API_KEY,
} as AuthOptions;

const httpLink = new HttpLink({ uri: url });

const link = ApolloLink.from([
  createAuthLink({ url, region: appsyncConfig.REGION, auth }),
  createSubscriptionHandshakeLink({ url, region, auth }, httpLink),
]);

export const appsyncClient = new ApolloClient({
  link,
  cache: new InMemoryCache(),
});
