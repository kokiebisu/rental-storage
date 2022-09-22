import { useNavigation } from "@react-navigation/native";

export const useFindMapScreen = () => {
  const navigation = useNavigation();
  const places = [
    {
      id: "sdfiojsdiojff2oij",
      title: "good garage 1",
      latitude: 28.32,
      longitude: -16.4,
      price: 25,
      emailAddress: "host1@gmail.com",
      available: "suitcase",
    },
    {
      id: "asdfsdavsd",
      title: "good garage 2",
      latitude: 28.3279822,
      longitude: -16.5124847,
      price: 35,
      emailAddress: "host2@gmail.com",
      available: "bag",
    },
    {
      id: "hhetrhe",
      title: "good garage 3",
      latitude: 28.2,
      longitude: -16.5124847,
      price: 50,
      emailAddress: "host3@gmail.com",
      available: "suitcase",
    },
  ];

  const handleNavigationToHomeScreen = 

  return {
    places,
    handleNavigationToHomeScreen: () => navigation.goBack(),
    handleNavigationToDetailsScreen: () => navigation.navigate('Details', {
      
    })
  };
};
