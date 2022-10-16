import * as React from "react";

import { usePostHomeScreen } from "./logic";
import Template from "./template";

export const PostHomeScreen = () => {
  const logic = usePostHomeScreen();
  return <Template {...logic} />;
};
