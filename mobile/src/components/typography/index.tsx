import * as React from "react";
import { Text } from "react-native";

export const Typography = ({
  variant,
  children,
  color = "black",
  bold = false,
  centered = false,
}) => {
  switch (variant) {
    case "h1":
      return (
        <Text
          style={{
            fontSize: 48,
            color,
            ...(bold && { fontWeight: "bold" }),
            ...(centered && { textAlign: "center" }),
          }}
        >
          {children}
        </Text>
      );
    case "h2":
      return (
        <Text
          style={{
            fontSize: 36,
            color,
            ...(bold && { fontWeight: "bold" }),
            ...(centered && { textAlign: "center" }),
          }}
        >
          {children}
        </Text>
      );
    case "h3":
      return (
        <Text
          style={{
            fontSize: 24,
            color,
            ...(bold && { fontWeight: "bold" }),
            ...(centered && { textAlign: "center" }),
          }}
        >
          {children}
        </Text>
      );
    case "h4":
      return (
        <Text
          style={{
            fontSize: 18,
            ...(bold && { fontWeight: "bold" }),
            ...(centered && { textAlign: "center" }),
          }}
        >
          {children}
        </Text>
      );
    default:
      throw new Error("Invalid typography variant");
  }
};
