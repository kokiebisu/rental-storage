import { useState } from "react";
import { useNavigation } from "@react-navigation/native";

export const useHomeScreen = () => {
  const navigation = useNavigation();
  const [viewport, setViewport] = useState(null);

  const [open, setOpen] = useState(false);
  const [value, setValue] = useState(null);
  const [items, setItems] = useState([
    { label: "Suitcase", value: "suitcase" },
    { label: "Bag", value: "bag" },
  ]);

  const onPressSearch = () => {
    navigation.navigate("Map", {
      payload: {
        viewport,
        category: value,
      },
    });
  };
  return {
    open,
    items,
    setOpen,
    setValue,
    setItems,
    value,
    setViewport,
    onPressSearch,
  };
};
