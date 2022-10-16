import * as React from "react";
import { useContext } from "react";
import {
  Dimensions,
  StyleSheet,
  View,
  Text,
  Pressable,
  Image,
  TextInput,
  Button,
} from "react-native";
import { GooglePlacesAutocomplete } from "react-native-google-places-autocomplete";
import { SafeAreaView } from "react-native-safe-area-context";
import { AuthContext } from "../../../../context/auth";
import { googleConfig } from "../../../../env";

import { SuggestionRow } from "../../stack-find/home/suggestion-row";

import "./styles";

export default () => {
  const { signOut } = useContext(AuthContext);
  return (
    <SafeAreaView style={{ paddingHorizontal: 20 }}>
      <Button title="Sign out" onPress={() => signOut()} />
    </SafeAreaView>
  );
};

const styles = StyleSheet.create({
  map: {
    width: Dimensions.get("window").width,
    height: Dimensions.get("window").height,
  },
});
