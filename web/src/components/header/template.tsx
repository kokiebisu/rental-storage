import { Menu, Center, createStyles, rem } from "@mantine/core";
import { IconChevronDown } from "@tabler/icons-react";

import { Button, MenuDropdown } from "..";
import Link from "next/link";

interface HeaderTemplateProps {
  links: {
    link: string;
    label: string;
    links?: { link: string; label: string }[];
  }[];
  onSignInClicked: () => void;
  isAuthenticated: boolean;
  profilePic?: string;
  handleSignout: () => void;
  handleGuestDashboardRedirect: () => void;
  handleLenderDashboardRedirect: () => void;
}

const useStyles = createStyles((theme) => ({
  link: {
    display: "block",
    lineHeight: 1,
    padding: `${rem(8)} ${rem(12)}`,
    borderRadius: theme.radius.sm,
    textDecoration: "none",
    color:
      theme.colorScheme === "dark"
        ? theme.colors.dark[0]
        : theme.colors.gray[7],
    fontSize: theme.fontSizes.sm,
    fontWeight: 500,

    "&:hover": {
      backgroundColor:
        theme.colorScheme === "dark"
          ? theme.colors.dark[6]
          : theme.colors.gray[0],
    },
  },
}));

const HeaderTemplate = ({
  links,
  isAuthenticated,
  profilePic,
  onSignInClicked,
  handleSignout,
  handleGuestDashboardRedirect,
  handleLenderDashboardRedirect,
}: HeaderTemplateProps) => {
  const { classes } = useStyles();
  const items = links.map((link) => {
    const menuItems = link.links?.map((item) => (
      <Menu.Item key={item.link}>{item.label}</Menu.Item>
    ));

    if (menuItems) {
      return (
        <Menu
          key={link.label}
          trigger="hover"
          transitionProps={{ exitDuration: 0 }}
          withinPortal
        >
          <Menu.Target>
            <a
              href={link.link}
              className={classes.link}
              onClick={(event) => event.preventDefault()}
            >
              <Center>
                <span className="mr-1">{link.label}</span>
                <IconChevronDown size="0.9rem" stroke={1.5} />
              </Center>
            </a>
          </Menu.Target>
          <Menu.Dropdown>{menuItems}</Menu.Dropdown>
        </Menu>
      );
    }

    return (
      <a
        key={link.label}
        href={link.link}
        className={classes.link}
        onClick={(event) => event.preventDefault()}
      >
        {link.label}
      </a>
    );
  });

  return (
    <header className="border-b-gray-200 border-b md:px-0">
      <div className="max-w-8xl mx-auto">
        <div className="px-12 2xl:px-0">
          <div className="flex justify-between items-center my-3">
            <div>
              <Link href="/">Rental Storage</Link>
            </div>
            <div className="flex items-center">
              <div className="sm:none flex mx-2 mr-2">{items}</div>
              {isAuthenticated ? (
                <MenuDropdown
                  handleSignout={handleSignout}
                  profilePic={profilePic}
                  handleGuestDashboardRedirect={handleGuestDashboardRedirect}
                  handleLenderDashboardRedirect={handleLenderDashboardRedirect}
                />
              ) : (
                <Button onClick={onSignInClicked}>Sign In</Button>
              )}
            </div>
          </div>
          {/* Not adding now */}
          {/* <div className="flex justify-between items-center my-3">
            <div>Sorting Filters here...</div>
            <div className="flex items-center">
              <Button onClick={() => alert()}>Map Search</Button>
            </div>
          </div> */}
        </div>
      </div>
    </header>
  );
};

export default HeaderTemplate;
