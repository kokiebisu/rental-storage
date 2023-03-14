import { TextInput as MantineTextInput } from "@mantine/core";
import { ChangeEventHandler } from "react";

interface TextInputProps {
  label?: string;
  value: string;
  placeholder: string;
  onChange: ChangeEventHandler<HTMLInputElement>;
}

const TextInput = ({ label, value, placeholder, onChange }: TextInputProps) => {
  return (
    <MantineTextInput
      label={label}
      value={value}
      placeholder={placeholder}
      onChange={onChange}
    />
  );
};

export default TextInput;
