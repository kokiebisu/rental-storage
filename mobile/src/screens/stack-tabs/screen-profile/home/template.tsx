import * as React from "react";
import { useContext } from "react";
import { Dimensions, StyleSheet, View, Text, Button } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import { AuthContext } from "../../../../context/auth";

import "./styles";
import { ProfileContext } from "../../../../context/profile";

export default () => {
  const { signOut } = useContext(AuthContext);
  const { profileState } = useContext(ProfileContext);

  return (
    <SafeAreaView style={{ paddingHorizontal: 20 }}>
      <Button title="Sign out" onPress={() => signOut()} />
      <View>
        <Text>{profileState.emailAddress}</Text>
        <Text>
          {profileState.firstName} {profileState.lastName}
        </Text>
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
