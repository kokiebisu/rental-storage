import * as React from "react";
import { createNativeStackNavigator } from "@react-navigation/native-stack";

import { stacksConfig } from "../../../config";
import { BorrowerMapScreen } from "./map";
import { BorrowerHomeScreen } from "./home";
import { BorrowerDetailsScreen } from "./details";

const Stack = createNativeStackNavigator();

export const BorrowerStackScreen = () => {
  return (
    <Stack.Navigator screenOptions={stacksConfig}>
      <Stack.Screen name="Home" component={BorrowerHomeScreen} />
      <Stack.Screen name="Map" component={BorrowerMapScreen} />
      <Stack.Screen name="Details" component={BorrowerDetailsScreen} />
    </Stack.Navigator>
  );
};
