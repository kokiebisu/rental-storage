import { Text } from "@mantine/core";
import { Dropzone, IMAGE_MIME_TYPE, FileWithPath } from "@mantine/dropzone";

export default function ImageUploader({ uploadURL, handleImageChange }) {
  return (
    <Dropzone accept={IMAGE_MIME_TYPE} onDrop={handleImageChange}>
      <Text align="center">Drop images here</Text>
    </Dropzone>
  );
}
