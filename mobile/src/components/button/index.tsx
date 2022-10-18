import * as React from "react";
import { PrimaryButton } from "./button-primary";

export const Button = (props) => {
  switch (props.variant) {
    case "primary":
      return <PrimaryButton {...props} />;
    default:
      throw new Error("Invalid button variant");
  }
};
