import { Avatar as MantineAvatar } from "@mantine/core";

interface AvatarProps {
  imageUrl: string;
  radius: "xl";
}

const Avatar = ({ imageUrl, radius }: AvatarProps) => (
  <MantineAvatar src={imageUrl} radius={radius} />
);

export default Avatar;
