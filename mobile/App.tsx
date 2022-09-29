import * as React from "react";
import { StatusBar } from "expo-status-bar";
import { NavigationContainer } from "@react-navigation/native";

import { MyTabs } from "./src/navigation/tab";
import { ApolloProvider } from "@apollo/client";
import { appsyncClient } from "./src/integration/graphql";

import { ContextProvider } from "./src/context";

export default () => {
  return (
    <ApolloProvider client={appsyncClient}>
      <ContextProvider>
        <NavigationContainer>
          <MyTabs />
        </NavigationContainer>
        <StatusBar />
      </ContextProvider>
    </ApolloProvider>
  );
};
