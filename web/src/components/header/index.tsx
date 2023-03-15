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
  const { user } = useContext(AuthContext);

  return (
    <HeaderTemplate
      links={links}
      onSignInClicked={onSignInClicked}
      isAuthenticated={!!user}
    />
  );
};

export default Header;
