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
      const url = 'https://dev-rental-a-locker-listing-profile.s3.amazonaws.com/test.txt?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=ASIA5T7AZ3LEMVFXS26L%2F20221016%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20221016T023750Z&X-Amz-Expires=900&X-Amz-Security-Token=IQoJb3JpZ2luX2VjEBsaCXVzLWVhc3QtMSJGMEQCIGpBZ5enbc8tzUm5tAj8M9JvFrccv5zwYiwIdmFSRE%2FrAiAywJLfPpLAiahHYg%2FKyoHY0SjEZg1VhcxDiD2YN94m6irEAwjk%2F%2F%2F%2F%2F%2F%2F%2F%2F%2F8BEAAaDDkzNjIzNzQ1NjA3MiIMOpBDVw8o9PbZyvYlKpgDis%2Fd%2FpISmE%2BkrLtQN3Rgt0AGv8zVEL2WErBzCJyb6ngt3l4%2BppztiBBn27jPFtRZjcbWijrpr5QuFi0%2F3cz8HQWplGklL6I7Scn6094UQ1zmYdEpApvNo0gvaFQDkLQJcTwCUuwIt6x6bW%2B8MpCYQvGl2sW7ePsCn6Q0dOvt%2FW1X817zbw4t8w%2BMYb%2F%2Fg2Zz%2BSUiTS4J%2BgP2yIYybZy6A5bVw0oxVLgoUt%2Bs0QfaHtd06ycdlsRj0cgJtZQpFLHkzvZeddlm3iNjZMKe2Rww5RaL%2FR%2BpewKAPYtQcpZgfn6RQII%2FK42xC1FmBf9%2B4Wxcu3PZGU6en6UgGfJsoaaAhaQenVsUg2t1Ai1%2Bq1kFvpH%2F40D%2BhdKMmS0%2FTvTZS83%2BM4T6gw%2BNA6aOYYr%2F%2FEdR9GLZNw5Htbwl0kMFLQw9svXl%2Bil138D35quaoFOG4mM3wCDE0j5rULIqwKLemKls1Tftt1e4sDpxq617cpasEcQnXBaoCFxPmT67RNapz4j6W%2B5YVOFF4HGEiBWp%2F5Hu87EdSRDlAVs%2FMP7crZoGOp8BJvat6%2BWCY7NawxeO1UbM0AKDBqQt%2F2gbHd8TFbGNmoSrE6RJqI%2FgyO%2BA7ciOwgFpaLJlRrYhayIs4KIhbpTS2dnKxms4n6nrA25V0%2BRm1VARuGrJjYtU%2BjLJLgL51UfHJNpu4cuytApj6ckhCp3c8KWRalbOTuL4I2I%2FbnAxF3m6EUgqjJWgXodwXmGm9PxfeQA3aVwcwNnQjEGwoSHq&X-Amz-Signature=e9df714528ed6f1d9372cf956241531ed69b10db9e42301675528d5a249b596a&X-Amz-SignedHeaders=host%3Bx-amz-acl&x-amz-acl=public-read'
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
