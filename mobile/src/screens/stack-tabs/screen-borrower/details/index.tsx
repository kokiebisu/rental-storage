import * as React from "react";

import { useDetailsScreen } from "./logic";
import Template from "./template";

export const BorrowerDetailsScreen = () => {
  const logic = useDetailsScreen();
  return <Template {...logic} />;
};
