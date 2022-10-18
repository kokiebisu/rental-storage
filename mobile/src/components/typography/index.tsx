import * as React from "react";
import { Text } from "react-native";

export const Typography = ({ variant, children, bold = false }) => {
  switch (variant) {
    case "h1":
      return (
        <Text style={{ fontSize: 48, ...(bold && { fontWeight: "bold" }) }}>
          {children}
        </Text>
      );
    case "h2":
      return (
        <Text style={{ fontSize: 36, ...(bold && { fontWeight: "bold" }) }}>
          {children}
        </Text>
      );
    case "h3":
      return (
        <Text style={{ fontSize: 24, ...(bold && { fontWeight: "bold" }) }}>
          {children}
        </Text>
      );
    case "h4":
      return <Text style={{ fontSize: 18 }}>{children}</Text>;
    default:
      throw new Error("Invalid variant");
  }
};
