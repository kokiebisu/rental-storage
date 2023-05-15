import axios from "axios";
import Image from "next/image";
import { useState } from "react";

export default function ImageUploader() {
  const [file, setFile] = useState(null);
  const [presignedUrl, setPresignedUrl] = useState("");

  const handleFileChange = (e) => {
    setFile(e.currentTarget.files[0]);
  };

  const handleGetPresignedUrl = async (e) => {
    try {
      const response = await axios.post(
        `${process.env.NEXT_PUBLIC_APIGATEWAY_ENDPOINT}/images`,
        {
          filename: file.name,
        }
      );
      setPresignedUrl(response["data"]["presignedUrl"]);
    } catch (error) {
      console.error(error);
    }
  };

  const handleUpload = async () => {
    try {
      await axios.put(presignedUrl, file, {
        headers: { "Content-Type": file.type },
      });
      alert("Image uploaded successfully");
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <div>
      <input type="file" onChange={handleFileChange} />
      <button onClick={handleGetPresignedUrl}>Get Presigned URL</button>

      {presignedUrl && (
        <div>
          <Image
            width={500}
            height={500}
            src={URL.createObjectURL(file)}
            alt=""
          />

          <button onClick={handleUpload}>Upload</button>
        </div>
      )}
    </div>
  );
}
