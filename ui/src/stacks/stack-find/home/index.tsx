import * as React from "react";
import { useState } from "react";
import { useNavigation } from "@react-navigation/native";
import { Text, Pressable, View } from "react-native";
import { GooglePlacesAutocomplete } from "react-native-google-places-autocomplete";
import { SafeAreaView } from "react-native-safe-area-context";
import DropDownPicker from "react-native-dropdown-picker";

import { googleConfig } from "../../../env";
import styles from "./styles";
import { SuggestionRow } from "./suggestion-row";
import { GET_TODOS } from "../../../graphql/all-todo";
import { useQuery } from "@apollo/client";

export const FindHomeScreen = () => {
  const navigation = useNavigation();
  const [viewport, setViewport] = useState(null);

  const [open, setOpen] = useState(false);
  const [value, setValue] = useState(null);
  const [items, setItems] = useState([
    { label: "Suitcase", value: "suitcase" },
    { label: "Bag", value: "bag" },
  ]);

  const {
    loading: listLoading,
    data: listData,
    error: listError,
  } = useQuery(GET_TODOS);

  const onPressSearch = () => {
    navigation.navigate("Map", {
      payload: {
        viewport,
        category: value,
      },
    });
  };

  return (
    <SafeAreaView style={{ flex: 1, paddingHorizontal: 10, paddingTop: 30 }}>
      <View style={{ flexGrow: 1 }}>
        <View style={{ paddingHorizontal: 10 }}>
          <GooglePlacesAutocomplete
            placeholder="Find location"
            onPress={(data, details = null) => {
              setViewport(details.geometry.viewport);
            }}
            fetchDetails
            styles={{
              textInput: styles.textInput,
            }}
            query={{
              key: googleConfig.API_KEY,
              language: "en",
              types: "(cities)",
            }}
            suppressDefaultStyles
            renderRow={(item) => <SuggestionRow item={item} />}
          />
        </View>
        <View style={{ paddingHorizontal: 10 }}>
          <DropDownPicker
            open={open}
            value={value}
            items={items}
            setOpen={setOpen}
            setValue={setValue}
            setItems={setItems}
          />
        </View>
      </View>
      <View>
        {listData.getTodos.map((item) => {
          return (
            <View>
              <Text>{item.description}</Text>
            </View>
          );
        })}
      </View>

      <View style={{ padding: 10 }}>
        <Pressable
          style={{ padding: 16, backgroundColor: "black" }}
          onPress={onPressSearch}
        >
          <Text
            style={{
              color: "white",
              textAlign: "center",
              fontWeight: "bold",
              fontSize: 20,
            }}
          >
            Search
          </Text>
        </Pressable>
      </View>
    </SafeAreaView>
  );
};
