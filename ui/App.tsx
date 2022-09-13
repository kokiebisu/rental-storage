import { StatusBar } from "expo-status-bar";
import { NavigationContainer } from "@react-navigation/native";

import { MyTabs } from "./src/navigation/tab";

export default () => {
  return (
    <>
      <NavigationContainer>
        <MyTabs />
      </NavigationContainer>
      <StatusBar />
    </>
  );
};
