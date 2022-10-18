import * as React from "react";
import { useHomeScreen } from "./logic";
import Template from "./template";

export const BorrowerHomeScreen = () => {
  const logic = useHomeScreen();
  return <Template {...logic} />;
};
