import * as React from "react";
import { useState } from "react";
import { useNavigation } from "@react-navigation/native";
import { Text, Pressable, View } from "react-native";
import { GooglePlacesAutocomplete } from "react-native-google-places-autocomplete";
import { SafeAreaView } from "react-native-safe-area-context";

import { Secrets } from "../../../config/secrets";
import styles from "./styles";
import { SuggestionRow } from "./suggestion-row";

export const FindHomeScreen = () => {
  const navigation = useNavigation();
  const [viewport, setViewport] = useState(null);

  const onPressSearch = () => {
    navigation.navigate("Map", {
      viewport,
    });
  };

  return (
    <SafeAreaView style={{ flex: 1 }}>
      <View style={{ flexGrow: 1 }}>
        <View>
          <GooglePlacesAutocomplete
            placeholder="Find loacation"
            onPress={(data, details = null) => {
              console.log(data, details);

              setViewport(details.geometry.viewport);
            }}
            fetchDetails
            styles={{
              textInput: styles.textInput,
            }}
            query={{
              key: Secrets.GOOGLE_PLACES_API_KEY,
              language: "en",
              types: "(cities)",
            }}
            suppressDefaultStyles
            renderRow={(item) => <SuggestionRow item={item} />}
          />
        </View>
        <View style={{ backgroundColor: "red" }}></View>
      </View>
      <View style={{ padding: 10 }}>
        <Pressable
          style={{ padding: 10, backgroundColor: "black" }}
          onPress={onPressSearch}
        >
          <Text style={{ color: "white", textAlign: "center" }}>Search</Text>
        </Pressable>
      </View>
    </SafeAreaView>
  );
};
