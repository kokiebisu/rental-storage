import { useState } from "react";
import * as ImagePicker from "expo-image-picker";
import axios from "axios";

export const usePostHomeScreen = () => {
  const [image, setImage] = useState(null);
  const [title, setTitle] = useState("");
  const [price, setPrice] = useState("");
  const [latLng, setLatLng] = useState(null);

  const handlePostListing = async () => {
    // alert(
    //   JSON.stringify({
    //     image,
    //     title,
    //     price,
    //     latLng,
    //   })
    // );
    // if (!image || !title || !price || !latLng) {
    //   throw new Error('Field is missing')
    // }

    try {
      const url = 'https://dev-rental-a-locker-listing-profile.s3.amazonaws.com/test.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=ASIA5T7AZ3LEE7KD7YNK%2F20221016%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20221016T140208Z&X-Amz-Expires=3600&X-Amz-SignedHeaders=host%3Bx-amz-acl&X-Amz-Security-Token=IQoJb3JpZ2luX2VjECYaCXVzLWVhc3QtMSJGMEQCIB9AoR71Ncrli0z5eEjTMwMuywrOGWIIvED8aPjhw5zjAiA465AE4ShCUDQ9kAkdfvqfJja4Pl3Z0x7WqvzD7wqcPSrEAwjv%2F%2F%2F%2F%2F%2F%2F%2F%2F%2F8BEAAaDDkzNjIzNzQ1NjA3MiIMsATAAsIiws1pGaoNKpgDO5CPajDXJE0UnqXbVRY2p1bzIY5VfkAKdfLj6JnVFEVv2V71eajzpZis5PLSK33PcKoG0VC40HSLHz%2FvDCpU1gB6F%2B8%2BRAfAvJ4VqFG28m3fi6g%2FXLWz23v%2BzQiTlYHAq0g%2F8j6ATCaAxisxlQ7iK8b5Nzt8ERyq0avJU80qGpHABX8AifEx%2FizeAKEptBJZO32RZbMHvhhEYMGiMupgsg0lyFbueI92DHr74jD3FEHKwYAenjZAHp0vikf5i9wifXXYOhSsyo3Oe71zmeI64guBFcNPLIZERudtwMm87t6Nrbh0wi7d0NywuCWqjY%2Bl20s9EOco0q5QO9M3r%2BS0Jvqxnt0%2ByhyBtWLFxTAxlhcRAMrGPQmjQxfrLU62ybVzHp3Iju0mLskiTXVpOhVpaJG6%2BFdOm6O2oZNclZi4OxOih3YgcCQkP%2B9PnomEALKf%2BZaAtWYd8ZvxH5TFy1VMkRGks2UR3LoGdQkhhyUGGDhdbIwHbh851rWh2vzT5JxFCF5S3QfWQgta9A95aTOO7C2P31VusGvsMN%2BdsJoGOp8Byr6aLPJinJlhC5zIaHw2DcgnYyKGwNnTSDbdD2ZNOjrlHIgXY1Y3KPikaYJyMrEEGD0lpI6v8cZTHyVxtb%2BgiXvF%2F68A0%2F%2FbCA89GDvG7yg%2FCRzoEyWylt1nE2UN9nvXSU%2FFiuxnLUA0dLLDLpzi4bQzSom2SMbDPAifXew%2FYDm%2B3fYb2Id1NZr01MHZj%2F9Q0GCl2iV%2BszzGtiGkAmbY&X-Amz-Signature=ea7eb9fd255e8e731dcdef65080d2ed184b5f8d7218897d6a2bbb551ad2e21f4'
      await uploadPhotoToS3(image.base64, image.uri, url)
    } catch (err) {
      console.error(err)
    }
    alert("Thank you for posting!");
  };

  const uploadPhotoToS3 = async (base64Data, uri, apiUrl) => {
    const uriParts = uri.split(".");
    const fileType = uriParts[uriParts.length - 1];
    const formData = new FormData()
    formData.append('photo', {
      uri, 
      name: `photo.${fileType}`,
      type: `image/${fileType}`
    } as any)
    
    // const uriParts = uri.split(".");
    // const fileType = uriParts[uriParts.length - 1];
    // const formData = new FormData()
    // formData.append('photo', {
    //   uri,
    //   name: `photo.${fileType}`,
    //   type: `image/${fileType}`
    // }as any)
    const config = {
      headers: {
        "x-amz-acl": "public-read",
        "Content-Encoding": "base64",
        "Content-Type": "image/jpeg"
      }
    }
    // console.log("FORMDATA: ", formData)
   
    try {
      // const response = await axios.post(url, formData, config)
      const response = await axios.put(apiUrl, Buffer.from(base64Data, "base64"), config)
      console.log("RESPONSE: ", response)
      return response;
    } catch (err) {
      console.error("ERROR: ", err)
    }
  }

  const handleImagePick = async () => {
    // No permissions request is necessary for launching the image library
    try {
      const result = await ImagePicker.launchImageLibraryAsync({
        mediaTypes: ImagePicker.MediaTypeOptions.All,
        allowsEditing: true,
        aspect: [4, 3],
        quality: 1,
        base64: true
      });
  
      if (!result.cancelled) {
        setImage(result);
      }
    } catch (err) {
      console.error(err)
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
