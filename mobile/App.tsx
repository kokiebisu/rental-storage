import * as React from "react";
import { StatusBar } from "expo-status-bar";
import { NavigationContainer } from "@react-navigation/native";
import { createNativeStackNavigator } from "@react-navigation/native-stack";
import { ApolloProvider } from "@apollo/client";

import { appsyncClient } from "./src/integration/graphql";

import { AuthContext } from "./src/context/auth";

import { AuthSignInScreen } from "./src/screens/stack-auth/screen-signin";
import { ContextProvider } from "./src/context";
import { SplashScreen, Tabs } from "./src/screens";
import { AuthSignUpScreen } from "./src/screens/stack-auth/screen-signup";

const Stack = createNativeStackNavigator();
export default () => (
  <ApolloProvider client={appsyncClient}>
    <ContextProvider>
      <Main />
    </ContextProvider>
    <StatusBar />
  </ApolloProvider>
);

const Main = () => {
  const { authState } = React.useContext(AuthContext);
  const isSignedIn = authState.userToken == null;
  return (
    <NavigationContainer>
      <Stack.Navigator>
        <Stack.Group>
          {isSignedIn ? (
            <Stack.Group>
              <Stack.Screen name="SignIn" component={AuthSignInScreen} />
              <Stack.Screen name="SignUp" component={AuthSignUpScreen} />
            </Stack.Group>
          ) : (
            <Stack.Group>
              <Stack.Screen
                name="Tabs"
                component={Tabs}
                options={{ headerShown: false }}
              />
            </Stack.Group>
          )}
          <Stack.Group>
            <Stack.Screen name="Splash" component={SplashScreen} />
          </Stack.Group>
        </Stack.Group>
      </Stack.Navigator>
    </NavigationContainer>
  );
};
