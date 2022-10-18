import * as React from "react";
import { View, Text } from "react-native";
import styles from "./styles";

export const SuggestionRow = ({ item }) => {
  return (
    <View style={styles.row}>
      <Text style={styles.locationText}>{item.description}</Text>
    </View>
  );
};
