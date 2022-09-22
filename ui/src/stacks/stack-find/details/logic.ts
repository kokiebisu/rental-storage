import { useRoute } from "@react-navigation/native";
import { Linking } from "react-native";

export const useDetailsScreen = () => {
  const { params } = useRoute();

  const listingData = {
    id: "1",
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
