import * as React from "react";
import { StatusBar } from "expo-status-bar";
import { NavigationContainer } from "@react-navigation/native";

import { ApolloProvider } from "@apollo/client";
import { appsyncClient } from "./src/integration/graphql";
import { createNativeStackNavigator } from "@react-navigation/native-stack";

import * as SecureStore from "expo-secure-store";
import { Tabs } from "./src/navigation/tab";
import { AuthContext } from "./src/context/auth";
import SplashScreen from "./src/stacks/screen-splash";
import { AuthSignInScreen } from "./src/stacks/stack-auth/signin";

const Stack = createNativeStackNavigator();

export default () => {
  const [state, dispatch] = React.useReducer(
    (prevState, action) => {
      switch (action.type) {
        case "RESTORE_TOKEN":
          return {
            ...prevState,
            userToken: action.token,
            isLoading: false,
          };
        case "SIGN_IN":
          return {
            ...prevState,
            isSignout: false,
            userToken: action.token,
          };
        case "SIGN_OUT":
          return {
            ...prevState,
            isSignout: true,
            userToken: null,
          };
      }
    },
    {
      isLoading: true,
      isSignout: false,
      userToken: null,
    }
  );

  React.useEffect(() => {
    // fetch the token from storage then navigate to our appropriate place
    const bootstrapAsync = async () => {
      let userToken;

      try {
        userToken = await SecureStore.getItemAsync("userToken");
        console.log("USER TOKEN: ", userToken);
      } catch (err) {
        // Restoring token failed
      }

      // after restoring token, we may need to validate it in production apps

      // this will switch to the App screen or Auth screen and this loading
      // screen will be unmounted and thrown away
      dispatch({ type: "RESTORE_TOKEN", token: userToken });
    };

    bootstrapAsync();
  }, []);

  const authContext = React.useMemo(
    () => ({
      signIn: async (data) => {
        // in a production app, we need to send some data (usually username, password) to server and get a token
        // we will also need to handle errors if sign in faield
        // after getting token, we need to persist the token using 'SecureStore'
        // in the example, we'll use a dummy token
        dispatch({ type: "SIGN_IN", token: "TOKEN_HERE" });
      },
      signOut: () => dispatch({ type: "SIGN_OUT" }),
      signUp: async (data) => {
        // in a production app, we need to send user data to server and get a token
        // we will also need to handle errors is signup failed
        // after getting token, we need to persist the token using 'SecureStore'
        // in the example, we'll use a dummy token
        dispatch({ type: "SIGN_IN", token: "TOKEN_HERE" });
      },
    }),
    []
  );

  const isSignedIn = state.userToken == null;

  return (
    <ApolloProvider client={appsyncClient}>
      <AuthContext.Provider value={authContext}>
        <NavigationContainer>
          <Stack.Navigator>
            <>
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
            </>
          </Stack.Navigator>
        </NavigationContainer>
      </AuthContext.Provider>
      <StatusBar />
    </ApolloProvider>
  );
};
