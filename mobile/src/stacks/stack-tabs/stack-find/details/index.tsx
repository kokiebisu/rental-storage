import * as React from "react";

import { useDetailsScreen } from "./logic";
import Template from "./template";

export const FindDetailsScreen = () => {
  const logic = useDetailsScreen();
  return <Template {...logic} />;
};
