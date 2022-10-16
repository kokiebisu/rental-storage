import * as React from "react";
import { useContext } from "react";
import { useQuery } from "@apollo/client";
import { Dimensions, StyleSheet, View, Text, Button } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import { AuthContext } from "../../../../context/auth";
import { QUERY_FIND_ME } from "../../../../graphql";

import "./styles";

export default () => {
  const { signOut } = useContext(AuthContext);
  const { data } = useQuery<any>(QUERY_FIND_ME, {
    fetchPolicy: "network-only",
  });
  if (!data) {
    alert("something went wrong");
  }
  const {
    findMe: { firstName, lastName, emailAddress, createdAt },
  } = data;
  return (
    <SafeAreaView style={{ paddingHorizontal: 20 }}>
      <Button title="Sign out" onPress={() => signOut()} />
      <View>
        <Text>
          {JSON.stringify({ firstName, lastName, emailAddress, createdAt })}
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
