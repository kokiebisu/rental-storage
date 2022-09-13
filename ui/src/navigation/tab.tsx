import { createBottomTabNavigator } from "@react-navigation/bottom-tabs";
import { tabsConfig } from "../config";
import { FindStackScreen, MyStorageStackScreen } from "../stacks";

const Tab = createBottomTabNavigator();

export const MyTabs = () => {
  return (
    <Tab.Navigator screenOptions={tabsConfig}>
      <Tab.Screen name="FindStack" component={FindStackScreen} />
      <Tab.Screen name="MyStorageStack" component={MyStorageStackScreen} />
    </Tab.Navigator>
  );
};
