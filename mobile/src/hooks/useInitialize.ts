import {
  ApolloClient,
  ApolloLink,
  HttpLink,
  InMemoryCache,
} from "@apollo/client";
import { AuthOptions, createAuthLink } from "aws-appsync-auth-link";
import { createSubscriptionHandshakeLink } from "aws-appsync-subscription-link";
import { useEffect, useState } from "react";
import { appsyncConfig } from "../env";
import * as SecureStore from "expo-secure-store";

const getUserToken = async () => {
  try {
    const token = await SecureStore.getItemAsync("userToken");
    console.debug("TOKEN: ", token);
    return token;
  } catch (err) {
    console.error(err);
  }
};

export const useInitialize = () => {
  const [client, setClient] = useState<any>(null);
  const [token, setToken] = useState(null);
  useEffect(() => {
    (async () => {
      const token = await getUserToken();
      setToken(token);
      const url = appsyncConfig.GRAPHQL_ENDPOINT;
      const region = appsyncConfig.REGION;
      const auth = {
        type: appsyncConfig.AUTHENTICATION_TYPE,
        token: token,
      } as AuthOptions;

      const httpLink = new HttpLink({ uri: url });

      const link = ApolloLink.from([
        createAuthLink({ url, region: appsyncConfig.REGION, auth }),
        createSubscriptionHandshakeLink({ url, region, auth }, httpLink),
      ]);

      setClient(
        new ApolloClient({
          link,
          cache: new InMemoryCache(),
        })
      );
    })();
  }, []);
  return { client };
};
