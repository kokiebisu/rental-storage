import * as React from "react";
import { Pressable, View } from "react-native";
import { Spacing } from "../../spacing";
import { Typography } from "../../typography";

export const PrimaryButton = ({
  onPress,
  label,
  backgroundColor = "#1D77AF",
}) => {
  return (
    <Pressable onPress={onPress}>
      <View
        style={{
          width: "100%",
          backgroundColor,
          paddingVertical: 15,
          borderRadius: 3,
        }}
      >
        <Typography color="white" centered variant="h3" bold>
          {label}
        </Typography>
      </View>
    </Pressable>
  );
};
