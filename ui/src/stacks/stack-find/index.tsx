import * as React from "react";
import { createNativeStackNavigator } from "@react-navigation/native-stack";

import { FindHomeScreen } from "./home";
import { stacksConfig } from "../../config";

const FindStack = createNativeStackNavigator();

export const FindStackScreen = () => {
  return (
    <FindStack.Navigator screenOptions={stacksConfig}>
      <FindStack.Screen name="Home" component={FindHomeScreen} />
    </FindStack.Navigator>
  );
};
