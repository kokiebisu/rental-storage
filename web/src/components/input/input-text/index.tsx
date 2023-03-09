import { TextInput as MantineTextInput } from "@mantine/core";

interface TextInputProps {
  label?: string;
  value: string;
  placeholder: string;
  onChange: () => void;
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
