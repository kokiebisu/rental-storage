import { Button as MantineButton } from "@mantine/core";

export type ButtonProps = {
  onClick?: () => void;
  type?: "button" | "reset" | "submit" | undefined;
  size?: string;
  children: React.ReactNode;
};

const Button = ({ type, size, onClick, children }: ButtonProps) => (
  <MantineButton size={size} type={type} variant="outline" onClick={onClick}>
    {children}
  </MantineButton>
);

export default Button;
