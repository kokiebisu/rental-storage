import { createStyles, Textarea, Group } from "@mantine/core";
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

const ContactUs = () => {
  const { classes } = useStyles();

  return (
    <div className={classes.form}>
      <TextInput
        label="Email"
        placeholder="your@email.com"
        required
        classNames={{ input: classes.input, label: classes.inputLabel }}
      />
      <TextInput
        label="Name"
        placeholder="John Doe"
        mt="md"
        classNames={{ input: classes.input, label: classes.inputLabel }}
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
        <Button className={classes.control} label="Send a Booking Request" />
      </Group>
    </div>
  );
};

export default ContactUs;
