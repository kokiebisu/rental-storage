import { useNavigation } from "@react-navigation/native";
import React, { useContext, useState } from "react";
import { Button, TextInput, View } from "react-native";
import { AuthContext } from "../../../context/auth";

export const AuthSignUpScreen = () => {
  const navigation = useNavigation();
  const [emailAddress, setEmailAddress] = useState("");
  const [password, setPassword] = useState("");
  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");

  const { signUp } = useContext(AuthContext);

  return (
    <View>
      <TextInput
        placeholder="Email Address"
        value={emailAddress}
        onChangeText={setEmailAddress}
      />
      <TextInput
        placeholder="Password"
        value={password}
        onChangeText={setPassword}
        secureTextEntry
      />
      <TextInput
        placeholder="First name"
        value={firstName}
        onChangeText={setFirstName}
        secureTextEntry
      />
      <TextInput
        placeholder="Last name"
        value={lastName}
        onChangeText={setLastName}
        secureTextEntry
      />
      <Button
        title="Sign up"
        onPress={() => signUp({ emailAddress, password, firstName, lastName })}
      />
      <Button title="Sign in" onPress={() => navigation.navigate("SignIn")} />
    </View>
  );
};
