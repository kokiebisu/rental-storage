import * as React from "react";
import { useFindMapScreen } from "./logic";

import Template from "./template";

export const FindMapScreen = () => {
  const logic = useFindMapScreen();

  return <Template {...logic} />;
};
