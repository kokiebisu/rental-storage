import * as React from "react";
import { createNativeStackNavigator } from "@react-navigation/native-stack";

import { LenderHomeScreen } from "./home";
import { LenderPostScreen } from "./post";
import { stacksConfig } from "../../../config";

const Stack = createNativeStackNavigator();

export const LenderStackScreen = () => {
  return (
    <Stack.Navigator screenOptions={stacksConfig}>
      <Stack.Screen name="Home" component={LenderHomeScreen} />
      <Stack.Screen name="Post" component={LenderPostScreen} />
    </Stack.Navigator>
  );
};
