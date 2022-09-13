import * as React from "react";
import { createNativeStackNavigator } from "@react-navigation/native-stack";

import { MyStorageHomeScreen } from "./home";
import { stacksConfig } from "../../config";

const MyStorageStack = createNativeStackNavigator();

export const MyStorageStackScreen = () => {
  return (
    <MyStorageStack.Navigator screenOptions={stacksConfig}>
      <MyStorageStack.Screen name="Home" component={MyStorageHomeScreen} />
    </MyStorageStack.Navigator>
  );
};
