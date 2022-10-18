import { useNavigation } from "@react-navigation/native";
import React, { useContext, useState } from "react";
import { Button, TextInput, View } from "react-native";
import { AuthContext } from "../../../context/auth";

export const AuthSignInScreen = () => {
  const navigation = useNavigation();
  const [emailAddress, setEmailAddress] = useState("");
  const [password, setPassword] = useState("");

  const { signIn } = useContext(AuthContext);

  return (
    <View>
      <TextInput
        placeholder="emailAddress"
        value={emailAddress}
        onChangeText={setEmailAddress}
        autoCapitalize="none"
      />
      <TextInput
        placeholder="Password"
        value={password}
        onChangeText={setPassword}
        secureTextEntry
      />
      <Button
        title="Sign in"
        onPress={() => signIn({ emailAddress, password })}
      />
      <Button title="Sign up" onPress={() => navigation.navigate("SignUp")} />
    </View>
  );
};
