import * as React from "react";
import { useRoute } from "@react-navigation/native";
import { useEffect, useRef, useState } from "react";
import {
  FlatList,
  Pressable,
  Text,
  useWindowDimensions,
  View,
} from "react-native";
import MapView, { PROVIDER_GOOGLE } from "react-native-maps";

import Post from "../../../components/card";
import { CustomMarker } from "../../../components";

export default ({
  places,
  handleNavigationToHomeScreen,
  handleNavigationToDetailsScreen,
}) => {
  const [selectedPlace, setSelectedPlace] = useState(null);
  const mapViewRef = useRef();
  const flatListRef = useRef();
  const route = useRoute();

  const {
    params: { payload },
  } = route;

  useEffect(() => {
    if (selectedPlace) {
      const index = places.findIndex((place) => place.id === selectedPlace.id);
      if (flatListRef.current) {
        flatListRef.current.scrollToIndex({ index });
      }

      if (mapViewRef.current) {
        mapViewRef.current.animateCamera({
          center: {
            latitude: selectedPlace.latitude - 0.1,
            longitude: selectedPlace.longitude,
          },
        });
      }
    }
  }, [selectedPlace]);

  const width = useWindowDimensions().width;

  return (
    <View style={{ width: "100%", height: "100%" }}>
      <Pressable
        style={{
          position: "absolute",
          top: 50,
          width: "100%",
          paddingHorizontal: 10,
          paddingVertical: 5,
          zIndex: 5,
        }}
        onPress={handleNavigationToHomeScreen}
      >
        <View
          style={{
            flexDirection: "row",
            paddingHorizontal: 10,
          }}
        >
          <View
            style={{
              marginRight: 20,
              padding: 10,
              borderRadius: 5,
              backgroundColor: "white",
            }}
          >
            <Text style={{ fontWeight: "bold", fontSize: 20 }}>Back</Text>
          </View>
          <View
            style={{
              backgroundColor: "white",
              flexGrow: 1,
              flexDirection: "row",
              borderRadius: 5,
              padding: 10,
            }}
          >
            <Text style={{ fontWeight: "bold", fontSize: 20 }}>
              {payload.category}
            </Text>
          </View>
        </View>
      </Pressable>
      <MapView
        ref={mapViewRef}
        style={{ width: "100%", height: "100%", zIndex: 1 }}
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
        {places
          .filter((place) => place.available == payload.category)
          .map((place, index) => (
            <CustomMarker
              onPress={() => {
                setSelectedPlace(place);
              }}
              key={index}
              isSelected={selectedPlace && place.id === selectedPlace.id}
              latitude={place.latitude}
              longitude={place.longitude}
              price={place.price}
            />
          ))}
      </MapView>
      <View
        style={{
          position: "absolute",
          zIndex: 5,
          bottom: 10,
          left: 0,
          right: 0,
        }}
      >
        <FlatList
          ref={flatListRef}
          data={places}
          horizontal
          showsHorizontalScrollIndicator={false}
          snapToInterval={useWindowDimensions().width - 30}
          snapToAlignment={"center"}
          decelerationRate={"fast"}
          renderItem={({ item }) => (
            <Post
              item={item}
              category={payload.category}
              width={width}
              onPress={() => handleNavigationToDetailsScreen(item)}
            />
          )}
          ListFooterComponent={() => (
            <View
              style={{ width: 15, height: 50, backgroundColor: "transparent" }}
            />
          )}
          ListHeaderComponent={() => (
            <View
              style={{ width: 15, height: 50, backgroundColor: "transparent" }}
            />
          )}
        />
      </View>
    </View>
  );
};
