import * as React from "react";
import { useState } from "react";
import {
  Dimensions,
  StyleSheet,
  View,
  Text,
  Pressable,
  Image,
  TextInput,
} from "react-native";
import { GooglePlacesAutocomplete } from "react-native-google-places-autocomplete";
import { SafeAreaView } from "react-native-safe-area-context";
import * as ImagePicker from "expo-image-picker";

import { Secrets } from "../../../config/secrets";
import { SuggestionRow } from "../../stack-find/home/suggestion-row";

import "./styles";

export const PostHomeScreen = () => {
  const [image, setImage] = useState(null);
  const [title, setTitle] = useState("");
  const [price, setPrice] = useState("");

  const pickImage = async () => {
    // No permissions request is necessary for launching the image library
    let result = await ImagePicker.launchImageLibraryAsync({
      mediaTypes: ImagePicker.MediaTypeOptions.All,
      allowsEditing: true,
      aspect: [4, 3],
      quality: 1,
    });

    if (!result.cancelled) {
      setImage(result.uri);
    }
  };

  return (
    <SafeAreaView style={{ paddingHorizontal: 20 }}>
      <View>
        <Text style={{ fontSize: 32, fontWeight: "bold" }}>
          Post your place
        </Text>
      </View>
      <View style={{ marginTop: 20 }}>
        <Pressable onPress={pickImage}>
          {image ? (
            <Image
              source={{ uri: image }}
              style={{ width: "100%", height: 300 }}
            />
          ) : (
            <View
              style={{
                height: 300,
                width: "100%",
                backgroundColor: "grey",
                alignItems: "center",
                justifyContent: "center",
              }}
            >
              <Text style={{ fontSize: 24 }}>Click here to upload image</Text>
            </View>
          )}
        </Pressable>
        {/* <Button title="Pick an image from camera roll" onPress={pickImage} /> */}
      </View>
      <View style={{ marginTop: 20 }}>
        <Text>Title</Text>
        <TextInput
          style={{ backgroundColor: "white", padding: 10, marginTop: 5 }}
          onChangeText={(e) => setTitle(e)}
          value={title}
          placeholder="Clean garage"
        />
      </View>
      <View style={{ marginTop: 20 }}>
        <Text>Location</Text>
        <View style={{ backgroundColor: "white", padding: 10, marginTop: 5 }}>
          <GooglePlacesAutocomplete
            placeholder="Find location"
            fetchDetails
            query={{
              key: Secrets.GOOGLE_PLACES_API_KEY,
              language: "en",
              types: "address",
            }}
            suppressDefaultStyles
            renderRow={(item) => <SuggestionRow item={item} />}
          />
        </View>
      </View>
      <View style={{ marginTop: 20 }}>
        <Text>Price</Text>
        <View style={{ marginTop: 5 }}>
          <TextInput
            style={{ backgroundColor: "white", padding: 10, marginTop: 5 }}
            onChangeText={(e) => setPrice(e)}
            value={price}
            placeholder="100"
          />
        </View>
      </View>
      <View style={{ marginTop: 20 }}>
        <Pressable
          style={{ backgroundColor: "black", width: "100%", padding: 16 }}
        >
          <Text
            style={{
              textAlign: "center",
              color: "white",
              fontSize: 20,
              fontWeight: "bold",
            }}
          >
            Post
          </Text>
        </Pressable>
      </View>
    </SafeAreaView>
  );
};

const styles = StyleSheet.create({
  map: {
    width: Dimensions.get("window").width,
    height: Dimensions.get("window").height,
  },
});
