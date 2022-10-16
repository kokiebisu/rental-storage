import { createBottomTabNavigator } from "@react-navigation/bottom-tabs";

import { tabsConfig } from "../../config";
import { FindStackScreen } from "./screen-find";
import { PostStackScreen } from "./screen-post";
import { ProfileStackScreen } from "./screen-profile";

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
        options={generateTabConfig("Profile")}
        name="ProfileStack"
        component={ProfileStackScreen}
      />
    </Tab.Navigator>
  );
};
