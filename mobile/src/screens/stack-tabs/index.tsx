import { createBottomTabNavigator } from "@react-navigation/bottom-tabs";

import { tabsConfig } from "../../config";
import { BorrowerStackScreen } from "./screen-borrower";
import { LenderStackScreen } from "./screen-lender";
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
        options={generateTabConfig("Borrower")}
        name="BorrowerStack"
        component={BorrowerStackScreen}
      />
      <Tab.Screen
        options={generateTabConfig("Lender")}
        name="LenderStack"
        component={LenderStackScreen}
      />
      <Tab.Screen
        options={generateTabConfig("Profile")}
        name="ProfileStack"
        component={ProfileStackScreen}
      />
    </Tab.Navigator>
  );
};
