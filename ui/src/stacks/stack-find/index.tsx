import * as React from "react";
import { createNativeStackNavigator } from "@react-navigation/native-stack";

import { stacksConfig } from "../../config";
import { FindMapScreen } from "./map";
import { FindHomeScreen } from "./home";

const FindStack = createNativeStackNavigator();

export const FindStackScreen = () => {
  return (
    <FindStack.Navigator screenOptions={stacksConfig}>
      <FindStack.Screen name="Home" component={FindHomeScreen} />
      <FindStack.Screen name="Map" component={FindMapScreen} />
    </FindStack.Navigator>
  );
};
