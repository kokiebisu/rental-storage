import {
  ApolloClient,
  ApolloLink,
  HttpLink,
  InMemoryCache,
  NormalizedCacheObject,
} from "@apollo/client";
import { AuthOptions, createAuthLink } from "aws-appsync-auth-link";
import { createSubscriptionHandshakeLink } from "aws-appsync-subscription-link";
import { appsyncConfig } from "../config";

const url = appsyncConfig.GRAPHQL_ENDPOINT;
const region = appsyncConfig.REGION;
const getAuthOption = (isAuthorized: boolean): AuthOptions => {
  if (isAuthorized) {
    return {
      type: "AWS_LAMBDA",
      token:
        (typeof window !== "undefined" &&
          localStorage.getItem("bearerToken")) ||
        "",
    };
  } else {
    return {
      type: "API_KEY",
      apiKey: appsyncConfig.API_KEY,
    };
  }
};

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
    link: ApolloLink.from([
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
    ]),
    cache: new InMemoryCache(),
    connectToDevTools: true,
  });
