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
import { appsyncConfig } from "../config";

export const CustomApolloProvider = ({ children }: any) => {
  const url = appsyncConfig.GRAPHQL_ENDPOINT;
  const region = appsyncConfig.REGION;
  const auth: AuthOptions = {
    type: "AWS_LAMBDA",
    token:
      typeof window !== "undefined"
        ? localStorage.getItem("authorizationToken") || ""
        : "",
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

  const link = ApolloLink.from([authLink, subscriptionLink]);

  link.request = (operation, forward) => {
    const token = localStorage.getItem("authorizationToken") || ""; // Get token from local storage
    operation.setContext({
      headers: {
        authorization: token ? `Bearer ${token}` : null,
      },
    });
    return forward ? forward(operation) : null;
  };

  const client: ApolloClient<NormalizedCacheObject> = new ApolloClient({
    link,
    cache: new InMemoryCache(),
    connectToDevTools: true,
  });

  return <ApolloProvider client={client}>{children}</ApolloProvider>;
};
