import { useEffect, useState } from "react";
import * as ImagePicker from "expo-image-picker";
import axios from "axios";
import { useLazyQuery } from "@apollo/client";
import { QUERY_GET_PRESIGNED_URL } from "../../../../graphql";

export const usePostHomeScreen = () => {
  const [image, setImage] = useState(null);
  const [title, setTitle] = useState("");
  const [price, setPrice] = useState("");
  const [latLng, setLatLng] = useState(null);
  const [url, setUrl] = useState(null);

  const [getImageUrls, { loading, error, data: imageUrlsData }] = useLazyQuery(
    QUERY_GET_PRESIGNED_URL
  );

  useEffect(() => {
    console.log("IMAGE URL: ", imageUrlsData);
    if (imageUrlsData) {
      const { url } = imageUrlsData.getPresignedURL;
      setUrl(url);
    }
  }, [imageUrlsData]);

  console.log("URL: ", url);

  const handlePostListing = async () => {
    try {
      await uploadPhotoToS3(image.base64, image.uri, url);
    } catch (err) {
      console.error(err);
    }
    alert("Thank you for posting!");
  };

  const uploadPhotoToS3 = async (base64Data, uri, apiUrl) => {
    const uriParts = uri.split(".");
    const fileType = uriParts[uriParts.length - 1];
    const formData = new FormData();
    formData.append("photo", {
      uri,
      name: `photo.${fileType}`,
      type: `image/${fileType}`,
    } as any);

    const config = {
      headers: {
        "x-amz-acl": "public-read",
        "Content-Encoding": "base64",
        "Content-Type": "image/jpeg",
      },
    };

    try {
      const response = await axios.put(
        apiUrl,
        Buffer.from(base64Data, "base64"),
        config
      );
      return response;
    } catch (err) {
      console.error("ERROR: ", err);
    }
  };

  const handleImagePick = async () => {
    // No permissions request is necessary for launching the image library
    try {
      const result = await ImagePicker.launchImageLibraryAsync({
        mediaTypes: ImagePicker.MediaTypeOptions.All,
        allowsEditing: true,
        aspect: [4, 3],
        quality: 1,
        base64: true,
      });

      if (!result.cancelled) {
        const arr = result.uri.split("/");
        const filename = arr[arr.length - 1];
        console.log("RESULT: ", filename);
        setImage(result);
        getImageUrls({
          variables: {
            filename,
          },
        });
      }
    } catch (err) {
      console.error(err);
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
