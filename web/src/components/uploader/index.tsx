import { Text } from "@mantine/core";
import { Dropzone, IMAGE_MIME_TYPE } from "@mantine/dropzone";

interface ImageUploaderProps {
  handleImageChange: (files: File[]) => void;
}

export default function ImageUploader({
  handleImageChange,
}: ImageUploaderProps) {
  return (
    <Dropzone accept={IMAGE_MIME_TYPE} onDrop={handleImageChange}>
      <Text align="center">Drop images here</Text>
    </Dropzone>
  );
}
