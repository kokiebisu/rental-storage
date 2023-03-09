import { Button as MantineButton } from "@mantine/core";

export interface ButtonProps {
  onClick: () => void;
  label: string;
}

const Button = ({ onClick, label }: ButtonProps) => (
  <MantineButton variant="outline" onClick={onClick}>
    {label}
  </MantineButton>
);

export default Button;
