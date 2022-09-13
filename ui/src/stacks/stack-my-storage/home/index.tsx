import * as React from "react";
import { Dimensions, StyleSheet, View } from "react-native";
import MapView from "react-native-maps";

export const MyStorageHomeScreen = () => {
  return (
    <View>
      <MapView
        style={styles.map}
        initialRegion={{
          latitude: 37.78825,
          longitude: -122.4324,
          latitudeDelta: 0.0922,
          longitudeDelta: 0.0421,
        }}
      />
    </View>
  );
};

const styles = StyleSheet.create({
  map: {
    width: Dimensions.get("window").width,
    height: Dimensions.get("window").height,
  },
});
