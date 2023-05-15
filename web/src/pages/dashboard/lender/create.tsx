import { Group, SimpleGrid, TextInput, Textarea, Title } from "@mantine/core";
import Button from "@/components/button";
import { useForm } from "@mantine/form";
import { ImageUploader } from "@/components";

export default function Dashboard() {
  const form = useForm({
    initialValues: {
      title: "",
      description: "",
      streetAddress: "",
      zip: "",
      city: "",
      province: "",
      country: "",
    },
    validate: {
      // title: (value) => value.trim().length < 2,
      // email: (value) => !/^\S+@\S+$/.test(value),
      // subject: (value) => value.trim().length === 0,
    },
  });

  const handleBookRequest = () => {
    alert(JSON.stringify(form.values));
  };

  return (
    <div className="py-12">
      <div className="max-w-3xl mx-auto px-5 2xl:px-0 w-full">
        <form onSubmit={form.onSubmit(() => {})}>
          <Title
            order={2}
            size="h1"
            sx={(theme) => ({
              fontFamily: `Greycliff CF, ${theme.fontFamily}`,
            })}
            weight={900}
            align="center"
          >
            Create a Listing
          </Title>
          <TextInput
            mt="md"
            label="Title"
            placeholder="Your title"
            name="title"
            variant="filled"
            {...form.getInputProps("title")}
          />
          <Textarea
            mt="md"
            label="Description"
            placeholder="Your description"
            maxRows={10}
            minRows={5}
            autosize
            name="description"
            variant="filled"
            {...form.getInputProps("description")}
          />
          <SimpleGrid
            cols={2}
            mt="md"
            breakpoints={[{ maxWidth: "sm", cols: 1 }]}
          >
            <TextInput
              label="Street Address"
              placeholder="Your street addresses"
              name="streetAddress"
              variant="filled"
              {...form.getInputProps("streetAddress")}
            />
            <TextInput
              label="Zip"
              placeholder="Zip"
              name="zip"
              variant="filled"
              {...form.getInputProps("zip")}
            />
          </SimpleGrid>
          <SimpleGrid
            cols={3}
            mt="md"
            breakpoints={[{ maxWidth: "sm", cols: 1 }]}
          >
            <TextInput
              label="City"
              placeholder="City"
              name="city"
              variant="filled"
              {...form.getInputProps("city")}
            />
            <TextInput
              label="Province"
              placeholder="Province"
              name="province"
              variant="filled"
              {...form.getInputProps("province")}
            />
            <TextInput
              label="Country"
              placeholder="Country"
              name="country"
              variant="filled"
              {...form.getInputProps("country")}
            />
          </SimpleGrid>
        </form>
        <ImageUploader />
        <div>Post the space image here</div>
        <Group position="center" mt="xl">
          <Button type="submit" size="md" onClick={handleBookRequest}>
            Send Book Request
          </Button>
        </Group>
      </div>
    </div>
  );
}
