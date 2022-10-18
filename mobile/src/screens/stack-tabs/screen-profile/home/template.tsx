import * as React from "react";
import { useContext } from "react";
import { View, Pressable } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import { AuthContext } from "../../../../context/auth";

import "./styles";
import { ProfileContext } from "../../../../context/profile";
import { Spacing } from "../../../../components/spacing";
import { Typography } from "../../../../components/typography";

export default () => {
  const { signOut } = useContext(AuthContext);
  const { profileState } = useContext(ProfileContext);

  return (
    <SafeAreaView style={{ paddingHorizontal: 20 }}>
      <View>
        <Spacing variant="lg">
          <Typography variant="h2">Profile</Typography>
        </Spacing>
      </View>
      <View style={{ flexDirection: "row", alignItems: "center" }}>
        <View style={{ marginRight: 16 }}>
          {/* Profile image */}
          <View
            style={{
              width: 48,
              height: 48,
              borderRadius: 100,
              backgroundColor: "black",
            }}
          />
        </View>
        <View>
          <Spacing variant="3sm">
            <Typography variant="h4">{profileState.emailAddress}</Typography>
          </Spacing>
          <Typography variant="h4">
            {profileState.firstName} {profileState.lastName}
          </Typography>
        </View>
      </View>
      <View>
        <Option label="Sign out" onPress={() => signOut()} />
      </View>
    </SafeAreaView>
  );
};

const Option = ({ label, onPress }) => {
  return (
    <Pressable onPress={onPress}>
      <Spacing variant="md">
        <Typography variant="h3">{label}</Typography>
      </Spacing>
    </Pressable>
  );
};
