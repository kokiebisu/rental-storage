import { useRoute } from "@react-navigation/native";
import * as React from "react";
import { Text, View } from "react-native";
import MapView, { PROVIDER_GOOGLE, Marker } from "react-native-maps";

export const FindMapScreen = () => {
  const route = useRoute();
  const {
    params: { payload },
  } = route;

  return (
    <View style={{ width: "100%", height: "100%" }}>
      <MapView
        style={{ width: "100%", height: "100%" }}
        initialRegion={{
          // latitude: payload.latLng.lat,
          // longitude: payload.latLng.lng,
          latitude: 28.3279822,
          longitude: -16.5124847,
          latitudeDelta: 0.8,
          longitudeDelta: 0.8,
        }}
        provider={PROVIDER_GOOGLE}
      >
        <Marker coordinate={{ latitude: 28.3279822, longitude: -16.5124847 }}>
          <View
            style={{
              padding: 5,
              backgroundColor: "white",
              borderRadius: 20,
              borderColor: "grey",
              borderWidth: 1,
            }}
          >
            <Text style={{ fontWeight: "700" }}>$100</Text>
          </View>
        </Marker>
      </MapView>
    </View>
  );
};
