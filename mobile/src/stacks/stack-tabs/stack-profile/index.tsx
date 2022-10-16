import * as React from "react";
import { createNativeStackNavigator } from "@react-navigation/native-stack";

import { ProfileHomeScreen } from "./home";
import { stacksConfig } from "../../../config";

const Stack = createNativeStackNavigator();

export const ProfileStackScreen = () => {
  return (
    <Stack.Navigator screenOptions={stacksConfig}>
      <Stack.Screen name="Home" component={ProfileHomeScreen} />
    </Stack.Navigator>
  );
};
