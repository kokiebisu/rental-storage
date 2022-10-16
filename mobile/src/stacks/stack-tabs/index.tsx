import { createBottomTabNavigator } from "@react-navigation/bottom-tabs";
import { createNativeStackNavigator } from "@react-navigation/native-stack";

import { tabsConfig } from "../../config";
import { FindStackScreen } from "./stack-find";
import { PostStackScreen } from "./stack-post";
import { ProfileStackScreen } from "./stack-profile";

const Stack = createNativeStackNavigator();
const Tab = createBottomTabNavigator();

const generateTabConfig = (title: string) => {
  return {
    title,
  };
};

export const Tabs = () => {
  return (
    <Tab.Navigator screenOptions={tabsConfig}>
      <Tab.Screen
        options={generateTabConfig("Find")}
        name="FindStack"
        component={FindStackScreen}
      />
      <Tab.Screen
        options={generateTabConfig("Post")}
        name="PostStack"
        component={PostStackScreen}
      />
      <Tab.Screen
        options={generateTabConfig("Post")}
        name="ProfileStack"
        component={ProfileStackScreen}
      />
    </Tab.Navigator>
  );
};
