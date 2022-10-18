import * as React from "react";
import { useBorrowerMapScreen } from "./logic";

import Template from "./template";

export const BorrowerMapScreen = () => {
  const logic = useBorrowerMapScreen();

  return <Template {...logic} />;
};
