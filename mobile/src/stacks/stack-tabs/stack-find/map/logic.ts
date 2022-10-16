import { useNavigation } from "@react-navigation/native";

export const useFindMapScreen = () => {
  const navigation = useNavigation();
  const places = [
    {
      id: "1",
      title: "good garage 1",
      latitude: 28.32,
      longitude: -16.4,
      price: 25,
      emailAddress: "lender1@gmail.com",
      available: "suitcase",
    },
    {
      id: "2",
      title: "good garage 2",
      latitude: 28.3279822,
      longitude: -16.5124847,
      price: 35,
      emailAddress: "lender2@gmail.com",
      available: "bag",
    },
    {
      id: "3",
      title: "good garage 3",
      latitude: 28.2,
      longitude: -16.5124847,
      price: 50,
      emailAddress: "lender3@gmail.com",
      available: "suitcase",
    },
  ];

  const handleNavigationToDetailsScreen = (item) => {
    navigation.navigate("Details", {
      payload: {
        listingId: item.id,
      },
    });
  };

  return {
    places,
    handleNavigationToHomeScreen: () => navigation.goBack(),
    handleNavigationToDetailsScreen,
  };
};
