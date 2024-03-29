import { Menu, UnstyledButton } from "@mantine/core";
import { forwardRef } from "react";
import { Avatar } from "../..";

interface MenuDropdownProps {
  profilePic?: string;
  handleSignout: () => void;
  handleGuestDashboardRedirect: () => void;
  handleLenderDashboardRedirect: () => void;
}

const MenuDropdown = ({
  profilePic,
  handleSignout,
  handleGuestDashboardRedirect,
  handleLenderDashboardRedirect,
}: MenuDropdownProps) => {
  return (
    <Menu>
      <Menu.Target>
        <AvatarClickable profilePic={profilePic} />
      </Menu.Target>
      <Menu.Dropdown right={2}>
        <Menu.Item onClick={handleGuestDashboardRedirect}>
          Guest Dashboard
        </Menu.Item>
        <Menu.Item onClick={handleLenderDashboardRedirect}>
          Lender Dashboard
        </Menu.Item>
        <Menu.Item onClick={handleSignout}>Sign out</Menu.Item>
      </Menu.Dropdown>
    </Menu>
  );
};

interface AvatarClickableProps
  extends React.ComponentPropsWithoutRef<"button"> {
  profilePic?: string;
}

const AvatarClickable = forwardRef<HTMLButtonElement, AvatarClickableProps>(
  ({ profilePic, ...others }: AvatarClickableProps, ref) => (
    <UnstyledButton
      ref={ref}
      sx={(theme) => ({
        display: "block",
        width: "100%",
        color:
          theme.colorScheme === "dark" ? theme.colors.dark[0] : theme.black,
      })}
      {...others}
    >
      <Avatar imageUrl={profilePic} radius="xl" />
    </UnstyledButton>
  )
);

AvatarClickable.displayName = "AvatarClickable";

export default MenuDropdown;
