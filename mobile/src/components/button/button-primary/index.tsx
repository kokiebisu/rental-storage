import * as React from "react";
import { Animated, Pressable } from "react-native";
import { Typography } from "../../typography";

export const PrimaryButton = ({
  onPress,
  label,
  backgroundColor = "#1D77AF",
}) => {
  const animated = new Animated.Value(1);

  const fadeIn = () => {
    Animated.timing(animated, {
      toValue: 0.1,
      duration: 100,
      useNativeDriver: true,
    }).start();
  };
  const fadeOut = () => {
    Animated.timing(animated, {
      toValue: 1,
      duration: 200,
      useNativeDriver: true,
    }).start();
  };

  return (
    <Pressable onPress={onPress} onPressIn={fadeIn} onPressOut={fadeOut}>
      <Animated.View
        style={{
          width: "100%",
          backgroundColor,
          paddingVertical: 15,
          borderRadius: 3,
          opacity: animated,
        }}
      >
        <Typography color="white" centered variant="h3" bold>
          {label}
        </Typography>
      </Animated.View>
    </Pressable>
  );
};
