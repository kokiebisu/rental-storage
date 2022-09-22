import * as React from "react";
import { Text, View, Pressable } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";

export default ({ listingData, handleOpenEmailComposer }) => {
  return (
    <SafeAreaView style={{ flex: 1, paddingHorizontal: 10, paddingTop: 30 }}>
      <View>
        <Text>yo</Text>
      </View>
      <View>
        <Pressable onPress={handleOpenEmailComposer}>
          <Text>Send message</Text>
        </Pressable>
      </View>
    </SafeAreaView>
  );
};
