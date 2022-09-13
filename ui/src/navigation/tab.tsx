import { createBottomTabNavigator } from "@react-navigation/bottom-tabs";
import { tabsConfig } from "../config";
import { FindStackScreen, MyStorageStackScreen } from "../stacks";

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
        options={generateTabConfig("Storage")}
        name="MyStorageStack"
        component={MyStorageStackScreen}
      />
    </Tab.Navigator>
  );
};
