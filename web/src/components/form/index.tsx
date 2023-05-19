import { createStyles, Textarea, Group } from "@mantine/core";
import { ChangeEventHandler } from "react";
import { Button, TextInput } from "..";

const useStyles = createStyles((theme) => ({
  form: {
    backgroundColor: theme.white,
    padding: theme.spacing.xl,
    borderRadius: theme.radius.md,
    boxShadow: theme.shadows.lg,
  },

  input: {
    backgroundColor: theme.white,
    borderColor: theme.colors.gray[4],
    color: theme.black,

    "&::placeholder": {
      color: theme.colors.gray[5],
    },
  },

  inputLabel: {
    color: theme.black,
  },

  control: {
    backgroundColor: theme.colors[theme.primaryColor][6],
  },
}));

interface BookingFormProps {
  email: string;
  name: string;
  onEmailChange: ChangeEventHandler<HTMLInputElement>;
  onNameChange: ChangeEventHandler<HTMLInputElement>;
}

const BookingForm = ({
  email,
  name,
  onEmailChange,
  onNameChange,
}: BookingFormProps) => {
  const { classes } = useStyles();

  return (
    <div className={classes.form}>
      <TextInput
        value={email}
        label="Email"
        onChange={onEmailChange}
        placeholder="your@email.com"
      />
      <TextInput
        value={name}
        onChange={onNameChange}
        label="Name"
        placeholder="John Doe"
      />
      <Textarea
        required
        label="Your message"
        placeholder="I want to order your goods"
        minRows={4}
        mt="md"
        classNames={{ input: classes.input, label: classes.inputLabel }}
      />

      <Group position="right" mt="md">
        <Button onClick={() => alert("should make a booking")}>
          Send a Booking Request
        </Button>
      </Group>
    </div>
  );
};

export default BookingForm;
