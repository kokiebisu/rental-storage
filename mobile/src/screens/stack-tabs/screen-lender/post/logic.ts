import { useContext, useEffect, useState } from "react";
import * as ImagePicker from "expo-image-picker";
import axios from "axios";
import { useLazyQuery, useMutation } from "@apollo/client";
import {
  MUTATION_ADD_LISTING,
  QUERY_GET_PRESIGNED_URL,
} from "../../../../graphql";
import { ProfileContext } from "../../../../context/profile";
import { Client } from "../../../../config/appsync";

export const useLenderPostScreen = () => {
  const [image, setImage] = useState(null);
  const [title, setTitle] = useState("");
  const [fee, setFee] = useState("");
  const [streetAddress, setStreetAddress] = useState("");
  const [latLng, setLatLng] = useState(null);
  const [url, setUrl] = useState(null);
  const { profile } = useContext(ProfileContext);

  const [getImageUrls, { loading, error, data: imageUrlsData }] = useLazyQuery(
    QUERY_GET_PRESIGNED_URL
  );
  const [
    addListing,
    {
      data: addListingData,
      loading: addListingLoading,
      error: addListingError,
    },
  ] = useMutation(MUTATION_ADD_LISTING, {
    client: Client,
  });

  useEffect(() => {
    if (imageUrlsData) {
      const { url } = imageUrlsData.getPresignedURL;
      setUrl(url);
    }
  }, [imageUrlsData]);

  const handleListingSubmit = async () => {
    try {
      await uploadPhotoToS3(image.base64, image.uri, url);
      await addListing({
        variables: {
          title,
          imageUrls: [url.split("?")[0]],
          lenderId: profile.uid,
          latitude: latLng.lat,
          longitude: latLng.lng,
          streetAddress,
          feeAmount: Number(fee),
          feeCurrency: "cad",
        },
      });
    } catch (err) {
      console.error(err);
    }
    alert("Thank you for Lendering!");
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

        setImage(result);
        getImageUrls({
          variables: {
            filename,
          },
          client: Client,
        });
      }
    } catch (err) {
      console.error(err);
    }
  };

  return {
    image,
    title,
    fee,
    handleTitleChange: setTitle,
    handlefeeChange: setFee,
    handleListingSubmit,
    handleSelectSuggestion: (data, details = null) => {
      setStreetAddress(data.description);
      setLatLng(details.geometry.location);
    },
    handleImagePick,
  };
};
