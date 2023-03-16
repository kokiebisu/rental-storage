import { useContext } from "react";

import { AuthContext } from "@/context/auth";
import HeaderTemplate from "./template";

interface HeaderSearchProps {
  onSignInClicked: () => void;
}

const Header = ({ onSignInClicked }: HeaderSearchProps) => {
  const links = [{ label: "Borrow", link: "/borrow" }];
  const { user, signout } = useContext(AuthContext);

  return (
    <HeaderTemplate
      links={links}
      onSignInClicked={onSignInClicked}
      isAuthenticated={!!user}
      handleSignout={signout}
    />
  );
};

export default Header;
