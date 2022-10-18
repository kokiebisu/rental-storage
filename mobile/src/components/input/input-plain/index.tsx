import * as React from "react";
import { TextInput } from "react-native";

export const PlainField = ({
  placeholder,
  value,
  onChangeText,
  capitalize = false,
  secure = false,
}) => {
  return (
    <TextInput
      style={{ fontSize: 24 }}
      placeholder={placeholder}
      value={value}
      onChangeText={onChangeText}
      secureTextEntry={secure}
      autoCapitalize={capitalize ? "words" : "none"}
    />
  );
};
