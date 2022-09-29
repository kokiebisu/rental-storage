import * as React from "react";
import { useHomeScreen } from "./logic";
import Template from "./template";

export const FindHomeScreen = () => {
  const logic = useHomeScreen();
  return <Template {...logic} />;
};
