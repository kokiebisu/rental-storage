import { createBottomTabNavigator } from "@react-navigation/bottom-tabs";
import { createNativeStackNavigator } from "@react-navigation/native-stack";
import { tabsConfig } from "../config";
import { FindStackScreen, PostStackScreen } from "../stacks";

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
        name="PostStock"
        component={PostStackScreen}
      />
    </Tab.Navigator>
  );
};

export const MainStack = () => {
  return (
    <Stack.Navigator>
      <Stack.Screen
        name="Tabs"
        component={Tabs}
        options={{ headerShown: false }}
      />
      {/* Profile */}
      {/* Settings */}
    </Stack.Navigator>
  );
};
