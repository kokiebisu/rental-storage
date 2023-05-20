import { useState } from "react";
import axios from "axios";
import { useForm } from "@mantine/form";
import {
  Group,
  SimpleGrid,
  TextInput,
  Textarea,
  Title,
  Image,
} from "@mantine/core";

import { ImageUploader, Button } from "@/components";
import { useMutation } from "@apollo/client";
import { SPACE_CREATE_MUTATION } from "@/graphql/mutations/space.graphql";
import { awsLambdaClient } from "@/clients";

export default function Dashboard() {
  const [selectedFile, setSelectedFile] = useState(null);
  const [createSpace, { error: spaceError }] = useMutation(
    SPACE_CREATE_MUTATION,
    {
      client: awsLambdaClient,
    }
  );

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

  const handleImageChange = async (files) => {
    setSelectedFile(files[0]);
  };

  const handleBookRequest = async (event) => {
    event.preventDefault();

    // check if all the input fields are filled
    const { title, description, streetAddress, zip, city, province, country } =
      form.values;

    if (
      !title ||
      !description ||
      !streetAddress ||
      !zip ||
      !city ||
      !province ||
      !country ||
      !selectedFile
    ) {
      alert("Please fill all the fields");
      return;
    }

    try {
      // get the presigned url
      const response = await axios.get(
        `${process.env.NEXT_PUBLIC_APIGATEWAY_ENDPOINT}/images?filename=${selectedFile.name}`
      );

      const uploadURL = response.data;
      // submitting image
      const formData = new FormData();
      Object.keys(uploadURL.fields).forEach((key) => {
        formData.append(key, uploadURL.fields[key]);
      });
      formData.append("file", selectedFile);

      await axios.post(uploadURL.url, formData);

      const input = {
        title,
        description,
        location: {
          address: streetAddress,
          city,
          country,
          countryCode: "CA",
          phone: "000-0000-0000",
          province,
          provinceCode: "BC",
          zip,
          coordinate: {
            latitude: 50,
            longitude: 50,
          },
        },
        imageUrls: [`${uploadURL.url}${selectedFile.name}`],
      };
      // create the space
      await createSpace({
        variables: input,
      });

      if (spaceError) {
        throw new Error(spaceError.message);
      }
      alert("Space created successfully!");
    } catch (error) {
      alert("Error creating space");
    }
  };

  return (
    <form onSubmit={handleBookRequest}>
      <div className="py-12">
        <div className="max-w-3xl mx-auto px-5 2xl:px-0 w-full">
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
          <div className="mt-6">
            <ImageUploader handleImageChange={handleImageChange} />
          </div>
          <SimpleGrid
            cols={4}
            breakpoints={[{ maxWidth: "sm", cols: 1 }]}
            mt={selectedFile ? "xl" : 0}
          >
            {selectedFile ? (
              <Image
                alt="uploaded image"
                src={URL.createObjectURL(selectedFile)}
              />
            ) : null}
          </SimpleGrid>
          <Group position="center" mt="xl">
            <Button type="submit" size="md">
              Post Space
            </Button>
          </Group>
        </div>
      </div>
    </form>
  );
}
