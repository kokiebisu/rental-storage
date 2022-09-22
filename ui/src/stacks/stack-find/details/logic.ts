import { Text, View, Linking, Pressable } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";

export const useDetailsScreen = () => {
  const listingData = {
    host: {
      emailAddress: "host@gmail.com",
    },
  };

  const handleOpenEmailComposer = async () => {
    await Linking.openURL(
      `mailto:${listingData.host.emailAddress}?subject=RentStorage&body=I want to store my ${category} at your place!`
    );
  };

  return {
    listingData,
    handleOpenEmailComposer,
  };
};
