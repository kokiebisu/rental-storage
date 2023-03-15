import { forwardRef } from "react";
import { IconChevronRight } from "@tabler/icons-react";
import { Group, Avatar, Text, UnstyledButton } from "@mantine/core";

interface UserButtonProps extends React.ComponentPropsWithoutRef<"button"> {
  image: string;
  name: string;
  email: string;
  icon?: React.ReactNode;
}

const MenuDropdown = forwardRef<HTMLButtonElement, UserButtonProps>(
  ({ image, name, email, icon, ...others }: UserButtonProps, ref) => (
    <UnstyledButton
      ref={ref}
      sx={(theme) => ({
        display: "block",
        width: "100%",
        padding: theme.spacing.md,
        color:
          theme.colorScheme === "dark" ? theme.colors.dark[0] : theme.black,

        "&:hover": {
          backgroundColor:
            theme.colorScheme === "dark"
              ? theme.colors.dark[8]
              : theme.colors.gray[0],
        },
      })}
      {...others}
    >
      <Group>
        <Avatar src={image} radius="xl" />
        {icon || <IconChevronRight size="1rem" />}
      </Group>
    </UnstyledButton>
  )
);

MenuDropdown.displayName = "MenuDropdown";

export default MenuDropdown;
