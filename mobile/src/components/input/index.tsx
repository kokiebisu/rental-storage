import * as React from "react";
import { PlainField } from "./input-plain";

export const Input = ({ variant, ...props }) => {
  switch (variant) {
    case "plain":
      return <PlainField {...props} />;
    default:
      throw new Error("Invalid input variant");
  }
};
