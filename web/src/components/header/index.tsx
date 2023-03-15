import { useContext, useEffect, useState } from "react";

import { AuthContext } from "@/context/auth";
import HeaderTemplate from "./template";

interface HeaderSearchProps {
  links: {
    link: string;
    label: string;
    links?: { link: string; label: string }[];
  }[];
  onSignInClicked: () => void;
}

const Header = ({ links, onSignInClicked }: HeaderSearchProps) => {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const { checkIsAuthenticated } = useContext(AuthContext);

  useEffect(() => {
    if (typeof window !== "undefined") {
      setIsAuthenticated(checkIsAuthenticated());
    }
  }, [checkIsAuthenticated]);

  return (
    <HeaderTemplate
      links={links}
      onSignInClicked={onSignInClicked}
      isAuthenticated={isAuthenticated}
    />
  );
};

export default Header;
