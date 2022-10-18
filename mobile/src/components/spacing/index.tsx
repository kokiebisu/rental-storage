import * as React from "react";
import { View } from "react-native";

export const Spacing = ({ variant, children }) => {
  switch (variant) {
    case "xl":
      return <View style={{ marginVertical: 24 }}>{children}</View>;
    case "lg":
      return <View style={{ marginVertical: 20 }}>{children}</View>;
    case "md":
      return <View style={{ marginVertical: 16 }}>{children}</View>;
    case "sm":
      return <View style={{ marginVertical: 12 }}>{children}</View>;
    default:
      throw new Error("Invalid spacing variant");
  }
};
