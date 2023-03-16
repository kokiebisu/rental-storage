import { Menu } from "@mantine/core";
import { Avatar } from "../..";

interface MenuDropdownProps {
  handleSignout: () => void;
}

const MenuDropdown = ({ handleSignout }: MenuDropdownProps) => {
  return (
    <Menu>
      <Menu.Target>
        <Avatar
          image="https://images.unsplash.com/photo-1508214751196-bcfd4ca60f91?ixid=MXwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHw%3D&ixlib=rb-1.2.1&auto=format&fit=crop&w=255&q=80"
          name="Harriette Spoonlicker"
          email="hspoonlicker@outlook.com"
        />
      </Menu.Target>
      <Menu.Dropdown>
        <Menu.Item>Profile</Menu.Item>
        <Menu.Item>Settings</Menu.Item>
        <Menu.Item onClick={handleSignout}>Sign out</Menu.Item>
      </Menu.Dropdown>
    </Menu>
  );
};

export default MenuDropdown;
