import {
  ApolloClient,
  ApolloLink,
  HttpLink,
  InMemoryCache,
  NormalizedCacheObject,
} from "@apollo/client";
import { setContext } from "@apollo/client/link/context";
import { AuthOptions, createAuthLink } from "aws-appsync-auth-link";
import { createSubscriptionHandshakeLink } from "aws-appsync-subscription-link";
import { appsyncConfig } from "../config";

const url = appsyncConfig.GRAPHQL_ENDPOINT;
const region = appsyncConfig.REGION;
const getAuthOption = (isAuthorized: boolean): AuthOptions => {
  if (isAuthorized) {
    return {
      type: "AWS_LAMBDA",
      token: getBearerToken(),
    };
  } else {
    return {
      type: "API_KEY",
      apiKey: appsyncConfig.API_KEY,
    };
  }
};

export const getBearerToken = () => {
  const token =
    typeof localStorage !== "undefined" && localStorage.getItem("bearerToken");
  return token || "";
};

const authLink = setContext((_, { headers }) => {
  // get the authentication token from local storage if it exists
  const token = localStorage.getItem("bearerToken");

  // return the headers to the context so httpLink can read them
  return {
    headers: {
      ...headers,
      authorization: token,
    },
  };
});

export const apiKeyClient: ApolloClient<NormalizedCacheObject> =
  new ApolloClient({
    link: ApolloLink.from([
      createAuthLink({
        url,
        region: appsyncConfig.REGION,
        auth: getAuthOption(false),
      }),
      createSubscriptionHandshakeLink(
        {
          url,
          region,
          auth: getAuthOption(false),
        },
        new HttpLink({ uri: url })
      ),
    ]),
    cache: new InMemoryCache(),
    connectToDevTools: true,
  });

export const awsLambdaClient: ApolloClient<NormalizedCacheObject> =
  new ApolloClient({
    link: authLink.concat(
      ApolloLink.from([
        createAuthLink({
          url,
          region: appsyncConfig.REGION,
          auth: getAuthOption(true),
        }),
        createSubscriptionHandshakeLink(
          {
            url,
            region,
            auth: getAuthOption(true),
          },
          new HttpLink({ uri: url })
        ),
      ])
    ),
    cache: new InMemoryCache(),
    connectToDevTools: true,
  });
