import { useLocalStorage } from "@/hooks/useLocalStorage";
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
import { createContext, useEffect, useState } from "react";
import { appsyncConfig } from "../config";

export const CustomApollo = createContext({
  client: new ApolloClient({
    link: ApolloLink.from([]),
    cache: new InMemoryCache(),
    connectToDevTools: true,
  }),
});

export const CustomApolloProvider = ({ children }: any) => {
  const { value } = useLocalStorage();
  const url = appsyncConfig.GRAPHQL_ENDPOINT;
  const region = appsyncConfig.REGION;
  console.log("VALUE: ", value);
  const auth: AuthOptions = {
    type: "AWS_LAMBDA",
    token:
      (typeof window !== "undefined" && localStorage.getItem("bearerToken")) ||
      "",
  };

  const authLink = createAuthLink({
    url,
    region: appsyncConfig.REGION,
    auth,
  });

  const subscriptionLink = createSubscriptionHandshakeLink(
    {
      url,
      region,
      auth,
    },
    new HttpLink({ uri: url })
  );

  const client: ApolloClient<NormalizedCacheObject> = new ApolloClient({
    link: ApolloLink.from([authLink, subscriptionLink]),
    cache: new InMemoryCache(),
    connectToDevTools: true,
  });

  return <ApolloProvider client={client}>{children}</ApolloProvider>;
};
