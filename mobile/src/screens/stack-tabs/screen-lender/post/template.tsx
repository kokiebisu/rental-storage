import * as React from "react";
import {
  Dimensions,
  StyleSheet,
  View,
  Text,
  Pressable,
  TextInput,
  Image,
} from "react-native";
import { GooglePlacesAutocomplete } from "react-native-google-places-autocomplete";
import { SafeAreaView } from "react-native-safe-area-context";
import { googleConfig } from "../../../../env";
import { SuggestionRow } from "../../screen-borrower/home/suggestion-row";

import "./styles";

export default ({
  image,
  title,
  price,
  handleTitleChange,
  handlePriceChange,
  handleListingSubmit,
  handleSelectSuggestion,
  handleImagePick,
}) => {
  console.log("IMAGE: ", image);
  return (
    <SafeAreaView style={{ paddingHorizontal: 20 }}>
      <View>
        <Text style={{ fontSize: 32, fontWeight: "bold" }}>
          List your place
        </Text>
      </View>
      <View style={{ marginTop: 20 }}>
        <Pressable onPress={handleImagePick}>
          {image ? (
            <View style={{ width: "100%", height: 300 }}>
              <Image
                // source={{ uri: image }}
                resizeMode="cover"
                source={{ uri: image.uri }}
                style={{
                  flex: 1,
                  width: undefined,
                  height: undefined,
                  resizeMode: "cover",
                }}
              />
            </View>
          ) : (
            // <View>
            //   <Text>image here</Text>
            // </View>
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
          onChangeText={handleTitleChange}
          value={title}
          placeholder="Clean garage"
          autoFocus
        />
      </View>
      <View style={{ marginTop: 20 }}>
        <Text>Location</Text>
        <View style={{ backgroundColor: "white", padding: 10, marginTop: 5 }}>
          <GooglePlacesAutocomplete
            placeholder="Find location"
            fetchDetails
            onPress={handleSelectSuggestion}
            query={{
              key: googleConfig.API_KEY,
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
            autoFocus
            style={{ backgroundColor: "white", padding: 10, marginTop: 5 }}
            onChangeText={handlePriceChange}
            value={price}
            placeholder="100"
          />
        </View>
      </View>
      <View style={{ marginTop: 20 }}>
        <Pressable
          onPress={handleListingSubmit}
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
            Submit
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
