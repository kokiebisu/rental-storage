import { Stepper, Group, Code } from "@mantine/core";
import { Button, TextInput, PasswordInput } from "../..";

export interface SignInModalTemplateProps {
  usernameProps: any;
  passwordProps: any;
  nameProps: any;
  emailProps: any;
  active: number;
  nextStep: () => void;
  prevStep: () => void;
}

const SignInModalTemplate = ({
  usernameProps,
  passwordProps,
  nameProps,
  emailProps,
  active,
  nextStep,
  prevStep,
}: SignInModalTemplateProps) => {
  return (
    <>
      <Stepper active={active} breakpoint="sm">
        <Stepper.Step label="First step" description="Profile settings">
          <TextInput
            label="Username"
            placeholder="Username"
            {...usernameProps}
          />
          <div className="mt-2">
            <PasswordInput {...passwordProps} />
          </div>
        </Stepper.Step>

        <Stepper.Step label="Second step" description="Personal information">
          <TextInput label="Name" placeholder="Name" {...nameProps} />
          <TextInput label="Email" placeholder="Email" {...emailProps} />
        </Stepper.Step>

        <Stepper.Completed>
          Completed! Form values:
          <Code block mt="xl">
            {JSON.stringify({ usernameProps }, null, 2)}
          </Code>
        </Stepper.Completed>
      </Stepper>

      <Group position="right" mt="xl">
        {active !== 0 && <Button label="Back" onClick={prevStep} />}
        {active !== 3 && <Button label="Next step" onClick={nextStep} />}
      </Group>
    </>
  );
};

export default SignInModalTemplate;
