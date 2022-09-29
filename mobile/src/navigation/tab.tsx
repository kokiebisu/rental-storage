import { createBottomTabNavigator } from "@react-navigation/bottom-tabs";
import { tabsConfig } from "../config";
import { FindStackScreen, PostStackScreen } from "../stacks";

const Tab = createBottomTabNavigator();

const generateTabConfig = (title: string) => {
  return {
    title,
  };
};

export const MyTabs = () => {
  return (
    <Tab.Navigator screenOptions={tabsConfig}>
      <Tab.Screen
        options={generateTabConfig("Find")}
        name="FindStack"
        component={FindStackScreen}
      />
      <Tab.Screen
        options={generateTabConfig("Post")}
        name="PostStock"
        component={PostStackScreen}
      />
    </Tab.Navigator>
  );
};
