import { useState } from "react";
import * as ImagePicker from "expo-image-picker";

export const usePostHomeScreen = () => {
  const [image, setImage] = useState(null);
  const [title, setTitle] = useState("");
  const [price, setPrice] = useState("");
  const [latLng, setLatLng] = useState(null);

  const handlePostListing = () => {
    alert(
      JSON.stringify({
        image,
        title,
        price,
        latLng,
      })
    );
    alert("Thank you for posting!");
  };

  const handleImagePick = async () => {
    // No permissions request is necessary for launching the image library
    let result = await ImagePicker.launchImageLibraryAsync({
      mediaTypes: ImagePicker.MediaTypeOptions.All,
      allowsEditing: true,
      aspect: [4, 3],
      quality: 1,
    });

    if (!result.cancelled) {
      setImage(result.uri);
    }
  };

  return {
    image,
    title,
    price,
    handleTitleChange: (e) => setTitle(e),
    handlePriceChange: (e) => setPrice(e),
    handlePostListing,
    handleSelectSuggestion: (data, details = null) => {
      setLatLng(details.geometry.viewport);
    },
    handleImagePick,
  };
};
