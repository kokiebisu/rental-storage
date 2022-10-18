import * as React from "react";

import { useLenderPostScreen } from "./logic";
import Template from "./template";

export const LenderPostScreen = () => {
  const logic = useLenderPostScreen();
  return <Template {...logic} />;
};
