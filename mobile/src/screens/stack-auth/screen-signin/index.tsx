import { useNavigation } from "@react-navigation/native";
import React, { useContext, useState } from "react";
import { View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import { Button } from "../../../components/button";
import { Input } from "../../../components/input";
import { Spacing } from "../../../components/spacing";
import { Typography } from "../../../components/typography";
import { AuthContext } from "../../../context/auth";

export const AuthSignInScreen = () => {
  const navigation = useNavigation();
  const [emailAddress, setEmailAddress] = useState("");
  const [password, setPassword] = useState("");

  const { signIn } = useContext(AuthContext);

  return (
    <SafeAreaView style={{ paddingHorizontal: 20 }}>
      <View>
        <Spacing variant="lg">
          <Typography variant="h2">Sign in</Typography>
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
            secure
          />
        </Spacing>
      </Spacing>
      <Spacing variant="ssm">
        <Button
          variant="primary"
          label="Sign in"
          onPress={() => signIn({ emailAddress, password })}
        />
      </Spacing>
      <Button
        variant="primary"
        label="Register"
        backgroundColor="#71A2B6"
        onPress={() => navigation.navigate("SignUp")}
      />
    </SafeAreaView>
  );
};
