import * as React from "react";
import { createNativeStackNavigator } from "@react-navigation/native-stack";

import { stacksConfig } from "../../config";

const Stack = createNativeStackNavigator();

export const AuthStack = () => {
  return (
    <Stack.Navigator screenOptions={stacksConfig}>
      <Stack.Screen name="Home" component={FindHomeScreen} />
      <Stack.Screen name="Map" component={FindMapScreen} />
    </Stack.Navigator>
  );
};
