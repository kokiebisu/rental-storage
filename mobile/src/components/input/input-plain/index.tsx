import * as React from "react";
import { TextInput } from "react-native";

export const PlainField = ({ placeholder, value, onChangeText }) => {
  return (
    <TextInput
      style={{ fontSize: 24 }}
      placeholder={placeholder}
      value={value}
      onChangeText={onChangeText}
    />
  );
};
