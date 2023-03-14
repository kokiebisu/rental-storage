import { PasswordInput as MantinePasswordInput } from "@mantine/core";

export interface PasswordInputProps {
  value: string;
  onChange: () => void;
}

const PasswordInput = ({ value, onChange }: PasswordInputProps) => {
  return (
    <MantinePasswordInput
      placeholder="Password"
      label="Password"
      description="Password must include at least one letter, number and special character"
      onChange={onChange}
      withAsterisk
      value={value}
    />
  );
};

export default PasswordInput;
