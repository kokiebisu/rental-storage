import * as React from "react";
import { StatusBar } from "expo-status-bar";
import { NavigationContainer } from "@react-navigation/native";
import { createNativeStackNavigator } from "@react-navigation/native-stack";

import { AuthContext } from "./src/context/auth";
import { ContextProvider } from "./src/context";
import { SplashScreen } from "./src/screens";
import { AuthSignInScreen } from "./src/screens/stack-auth";
import { AuthSignUpScreen } from "./src/screens/stack-auth";
import { Tabs } from "./src/screens/stack-tabs";

export const App = () => {
  return (
    <ContextProvider>
      <Main />
      <StatusBar />
    </ContextProvider>
  );
};

const Main = () => {
  const { authState } = React.useContext(AuthContext);
  const isSignedIn = authState.userToken == null;
  const Stack = createNativeStackNavigator();
  return (
    <NavigationContainer>
      <Stack.Navigator>
        <Stack.Group screenOptions={{ headerShown: false }}>
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

export default App;
