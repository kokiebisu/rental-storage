import * as React from "react";
import { StatusBar } from "expo-status-bar";
import { NavigationContainer } from "@react-navigation/native";
import { createNativeStackNavigator } from "@react-navigation/native-stack";
import { ApolloProvider } from "@apollo/client";

import { appsyncClient } from "./src/integration/graphql";
import { Tabs } from "./src/navigation/tab";
import { AuthContext } from "./src/context/auth";
import SplashScreen from "./src/stacks/screen-splash";
import { AuthSignInScreen } from "./src/stacks/stack-auth/signin";
import { ContextProvider } from "./src/context";

const Stack = createNativeStackNavigator();
export default () => {
  return (
    <ApolloProvider client={appsyncClient}>
      <ContextProvider>
        <Main />
      </ContextProvider>
      <StatusBar />
    </ApolloProvider>
  );
};

const Main = () => {
  const { authState } = React.useContext(AuthContext);
  const isSignedIn = authState.userToken == null;
  return (
    <NavigationContainer>
      <Stack.Navigator>
        <React.Fragment>
          {isSignedIn ? (
            <React.Fragment>
              <Stack.Screen name="SignIn" component={AuthSignInScreen} />
            </React.Fragment>
          ) : (
            <React.Fragment>
              <Stack.Screen
                name="Tabs"
                component={Tabs}
                options={{ headerShown: false }}
              />
            </React.Fragment>
          )}
          <Stack.Group>
            <Stack.Screen name="Splash" component={SplashScreen} />
          </Stack.Group>
        </React.Fragment>
      </Stack.Navigator>
    </NavigationContainer>
  );
};
