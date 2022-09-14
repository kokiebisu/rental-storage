import * as React from "react";
import { createNativeStackNavigator } from "@react-navigation/native-stack";

import { PostHomeScreen } from "./home";
import { stacksConfig } from "../../config";

const PostStack = createNativeStackNavigator();

export const PostStackScreen = () => {
  return (
    <PostStack.Navigator screenOptions={stacksConfig}>
      <PostStack.Screen name="Home" component={PostHomeScreen} />
    </PostStack.Navigator>
  );
};
