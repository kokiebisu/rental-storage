import { useNavigation } from "@react-navigation/native";
import React, { useContext, useState } from "react";
import { View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import { Button } from "../../../components/button";
import { Input } from "../../../components/input";
import { Spacing } from "../../../components/spacing";
import { Typography } from "../../../components/typography";
import { AuthContext } from "../../../context/auth";

export const AuthSignUpScreen = () => {
  const navigation = useNavigation();
  const [emailAddress, setEmailAddress] = useState("");
  const [password, setPassword] = useState("");
  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");

  const { signUp } = useContext(AuthContext);

  return (
    <SafeAreaView style={{ paddingHorizontal: 20 }}>
      <View>
        <Spacing variant="lg">
          <Typography variant="h2">Sign up</Typography>
        </Spacing>
      </View>
      <Spacing variant="lg">
        <Spacing variant="md">
          <Input
            variant="plain"
            placeholder="Email Address"
            value={emailAddress}
            onChangeText={setEmailAddress}
          />
        </Spacing>
        <Spacing variant="md">
          <Input
            placeholder="Password"
            variant="plain"
            value={password}
            onChangeText={setPassword}
          />
        </Spacing>
        <Spacing variant="md">
          <Input
            variant="plain"
            placeholder="First name"
            value={firstName}
            onChangeText={setFirstName}
          />
        </Spacing>
        <Spacing variant="md">
          <Input
            variant="plain"
            placeholder="Last name"
            value={lastName}
            onChangeText={setLastName}
          />
        </Spacing>
      </Spacing>
      <Spacing variant="ssm">
        <Button
          variant="primary"
          label="Sign up"
          onPress={() =>
            signUp({ emailAddress, password, firstName, lastName })
          }
        />
      </Spacing>
      <Button
        variant="primary"
        label="Already a user"
        backgroundColor="#71A2B6"
        onPress={() => navigation.navigate("SignIn")}
      />
    </SafeAreaView>
  );
};
