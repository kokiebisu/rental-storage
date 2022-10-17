import { createContext, useContext } from "react";
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
import { AuthContext } from "./auth";

const getUserToken = async () => {
  try {
    const token = await SecureStore.getItemAsync("userToken");
    return token;
  } catch (err) {
    console.error(err);
  }
};

export const ApolloContext = createContext(null);

export const ApolloContextProvider = ({ children }) => {
  const [client, setClient] = useState<any>(null);
  const { authState } = useContext(AuthContext);

  useEffect(() => {
    if (!authState.userToken) {
      (async () => {
        console.log("TRIGGERED");
        const token = await getUserToken();
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
    }
  }, [authState.userToken]);

  return (
    <ApolloContext.Provider value={client}>{children}</ApolloContext.Provider>
  );
};
