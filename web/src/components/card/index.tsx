import {
  Card as MantineCard,
  Image as MantineImage,
  Text as MantineText,
  Badge as MantineBadge,
  Button as MantineButton,
  Group as MantineGroup,
} from "@mantine/core";

const Card = () => {
  return (
    <MantineCard shadow="sm" padding="lg" radius="md" withBorder>
      <MantineCard.Section>
        <MantineImage
          src="https://images.unsplash.com/photo-1527004013197-933c4bb611b3?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=720&q=80"
          height={160}
          alt="Norway"
        />
      </MantineCard.Section>

      <MantineGroup position="apart" mt="md" mb="xs">
        <MantineText weight={500}>Norway Fjord Adventures</MantineText>
        <MantineBadge color="pink" variant="light">
          On Sale
        </MantineBadge>
      </MantineGroup>

      <MantineText size="sm" color="dimmed">
        With Fjord Tours you can explore more of the magical fjord landscapes
        with tours and activities on and around the fjords of Norway
      </MantineText>

      <MantineButton variant="light" color="blue" fullWidth mt="md" radius="md">
        Book classic tour now
      </MantineButton>
    </MantineCard>
  );
};

export default Card;
