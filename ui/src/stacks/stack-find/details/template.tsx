import * as React from "react";
import { Text, View, Pressable, Image, ScrollView } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";

export default ({ listingData, handleOpenEmailComposer }) => {
  return (
    <SafeAreaView style={{ flexGrow: 1 }}>
      <ScrollView style={{ paddingHorizontal: 20 }}>
        <View
          style={{
            marginBottom: 10,
            maxHeight: 300,
            width: "100%",
          }}
        >
          <Image
            style={{
              width: "100%",
              aspectRatio: 1.4,
              resizeMode: "cover",
            }}
            source={{
              uri: "https://storage.googleapis.com/gweb-uniblog-publish-prod/images/Garage__wide_shot_2.max-1300x1300.png",
            }}
          />
        </View>
        <View style={{ marginBottom: 10 }}>
          <Text style={{ fontSize: 32 }}>Clean Garage</Text>
        </View>
        <View style={{ marginBottom: 20 }}>
          <Text style={{ fontSize: 24 }}>$150/month</Text>
        </View>
        <View>
          <Text style={{ fontSize: 16 }}>
            It is a clean garage and I maintain it well.
          </Text>
        </View>
      </ScrollView>
      <View
        style={{
          position: "absolute",
          width: "100%",
          bottom: 0,
          padding: 10,
        }}
      >
        <Pressable
          style={{ backgroundColor: "black", width: "100%", padding: 15 }}
          onPress={handleOpenEmailComposer}
        >
          <Text
            style={{
              textAlign: "center",
              color: "white",
              fontSize: 24,
              fontWeight: "700",
            }}
          >
            Send message
          </Text>
        </Pressable>
      </View>
    </SafeAreaView>
  );
};
