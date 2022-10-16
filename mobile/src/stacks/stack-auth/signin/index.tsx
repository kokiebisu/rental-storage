import { useNavigation } from "@react-navigation/native";
import React, { useContext, useState } from "react";
import { Button, TextInput, View } from "react-native";
import { AuthContext } from "../../../context/auth";

export const AuthSignInScreen = () => {
  const navigation = useNavigation();
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const { signIn } = useContext(AuthContext);

  return (
    <View>
      <TextInput
        placeholder="Username"
        value={username}
        onChangeText={setUsername}
      />
      <TextInput
        placeholder="Password"
        value={password}
        onChangeText={setPassword}
        secureTextEntry
      />
      <Button title="Sign in" onPress={() => signIn({ username, password })} />
      <Button title="Sign up" onPress={() => navigation.navigate("SignUp")} />
    </View>
  );
};
