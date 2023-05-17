import axios from "axios";
import Image from "next/image";
import { useState } from "react";

export default function ImageUploader() {
  const [selectedFile, setSelectedFile] = useState(null);
  const [uploadURL, setUploadURL] = useState(null);

  const handleImageChange = (event) => {
    setSelectedFile(event.target.files[0]);
  };

  const handleSubmit = async (event) => {
    event.preventDefault();
    const formData = new FormData();
    formData.append("file", selectedFile);
    Object.keys(uploadURL.fields).forEach((key) => {
      formData.append(key, uploadURL.fields[key]);
    });

    try {
      await axios.post(uploadURL.url, formData);
      alert("Image uploaded successfully!");
    } catch (error) {
      console.log(error);
      alert("Error uploading image");
    }
  };

  const fetchUploadURL = async () => {
    try {
      const response = await axios.get(
        `${process.env.NEXT_PUBLIC_APIGATEWAY_ENDPOINT}/images`
      );
      setUploadURL(response.data);
    } catch (error) {
      console.log(error);
    }
  };

  console.log("UPLOADED URL: ", uploadURL);
  console.log("SELECTED FILE: ", selectedFile);

  return (
    <div>
      <h3>Upload Image</h3>
      <form onSubmit={handleSubmit}>
        <input type="file" onChange={handleImageChange} />
        <br />
        <button disabled={!selectedFile} onClick={fetchUploadURL}>
          Get Upload URL
        </button>
        <br />
        <br />
        {uploadURL && (
          <div>
            Upload URL: <a href={uploadURL.url}>{uploadURL.url}</a>
            <br />
            <br />
            <button type="submit">Upload Image</button>
          </div>
        )}
      </form>
    </div>
  );
}
