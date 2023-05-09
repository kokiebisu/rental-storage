import { useMutation, useQuery } from "@apollo/client";
import { useRouter } from "next/router";

import { apiKeyClient, awsLambdaClient } from "@/clients";
import { Button, Grid } from "@/components";
import { DefaultLayout } from "@/layout";
import { FIND_SPACE_QUERY } from "@/graphql/queries";
import {
  ActionIcon,
  Group,
  TextInput,
  Textarea,
  createStyles,
  rem,
} from "@mantine/core";
import { useEffect, useState } from "react";
import { MAKE_BOOKING_MUTATION } from "@/graphql/mutations/booking.graphql";
import { Spinner } from "@/components/spinner";
import {
  IconBrandInstagram,
  IconBrandTwitter,
  IconBrandYoutube,
} from "@tabler/icons-react";
import { DatePicker } from "@mantine/dates";

const useStyles = createStyles((theme) => ({
  wrapper: {
    minHeight: 400,
    boxSizing: "border-box",
    backgroundImage: `linear-gradient(-60deg, ${
      theme.colors[theme.primaryColor][4]
    } 0%, ${theme.colors[theme.primaryColor][7]} 100%)`,
    borderRadius: theme.radius.md,
    padding: `calc(${theme.spacing.xl} * 2.5)`,

    [theme.fn.smallerThan("sm")]: {
      padding: `calc(${theme.spacing.xl} * 1.5)`,
    },
  },

  title: {
    fontFamily: `Greycliff CF, ${theme.fontFamily}`,
    color: theme.white,
    lineHeight: 1,
  },

  description: {
    color: theme.colors[theme.primaryColor][0],
    maxWidth: rem(300),

    [theme.fn.smallerThan("sm")]: {
      maxWidth: "100%",
    },
  },

  form: {
    backgroundColor: theme.white,
    padding: theme.spacing.xl,
    borderRadius: theme.radius.md,
    boxShadow: theme.shadows.lg,
  },

  social: {
    color: theme.white,

    "&:hover": {
      color: theme.colors[theme.primaryColor][1],
    },
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

const social = [IconBrandTwitter, IconBrandYoutube, IconBrandInstagram];

export default function SpaceDetailsPage() {
  const router = useRouter();
  const [message, setMessage] = useState("");
  const [date, setDate] = useState<[Date | null, Date | null]>([null, null]);
  const { classes } = useStyles();
  const { id } = router.query;
  const {
    data: spaceData,
    error: spaceError,
    loading: spaceLoading,
  } = useQuery(FIND_SPACE_QUERY, {
    variables: { id },
    client: apiKeyClient,
  });
  const [
    makeBooking,
    { data: bookingData, loading: bookingLoading, error: bookingError },
  ] = useMutation(MAKE_BOOKING_MUTATION, {
    client: awsLambdaClient,
  });

  useEffect(() => {
    if (!bookingLoading && !bookingError && bookingData) {
      router.push("/success");
    }
  }, [bookingLoading, bookingError, bookingData, router]);

  const handleMakeBooking = () => {
    makeBooking({
      variables: {
        description,
        endDate: "2021-09-01T00:00:00.000Z",
        startDate: "2021-09-01T00:00:00.000Z",
        imageUrls:
          "https://assets.awaytravel.com/spree/products/27598/original/PDP_PC_Sand_CAR_01.jpg",
        spaceId: id,
      },
    });
  };

  if (spaceLoading) {
    return (
      <div className="absolute top-1/2 left-1/2">
        <Spinner />
      </div>
    );
  }

  if (spaceError) {
    return <div>error</div>;
  }

  const icons = social.map((Icon, index) => (
    <ActionIcon
      key={index}
      size={28}
      className={classes.social}
      variant="transparent"
    >
      <Icon size="1.4rem" stroke={1.5} />
    </ActionIcon>
  ));

  const {} = spaceData;

  return (
    <DefaultLayout>
      <div className="mb-12">
        <Grid />
      </div>
      <div className="max-w-7xl mx-auto px-5 2xl:px-0 w-full">
        <div className="flex">
          <div className="flex-1 pr-12">
            <div className="mb-4">
              <h2 className="font-bold text-3xl">{spaceData.space.title}</h2>
            </div>
            <div className="mb-12">{spaceData.space.description}</div>
            <div>
              {" "}
              <Group position="center">
                <DatePicker
                  type="range"
                  allowSingleDateInRange
                  value={date}
                  onChange={setDate}
                />
              </Group>
            </div>
          </div>
          <div className="w-[450px]">
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
                placeholder="I want to book your place..."
                minRows={4}
                mt="md"
                onChange={(e) => setMessage(e.target.value)}
              />
              <Group position="right" mt="md">
                <Button
                  onClick={handleMakeBooking}
                  label="Make Booking Request"
                />
              </Group>
            </div>
          </div>
        </div>
      </div>
    </DefaultLayout>
  );
}
