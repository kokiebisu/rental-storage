import { useMutation, useQuery } from "@apollo/client";
import { useRouter } from "next/router";

import { apiKeyClient, awsLambdaClient } from "@/clients";
import { Button, Grid } from "@/components";
import { DefaultLayout } from "@/layout";
import { FIND_SPACE_QUERY } from "@/graphql/queries";
import { Textarea } from "@mantine/core";
import { useEffect, useState } from "react";
import { MAKE_BOOKING_MUTATION } from "@/graphql/mutations/booking.graphql";

export default function SpaceDetailsPage() {
  const router = useRouter();
  const [description, setDescription] = useState("");
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
    return <div>loading</div>;
  }

  if (spaceError) {
    return <div>error</div>;
  }

  return (
    <DefaultLayout>
      <div className="my-6">
        <Grid />
      </div>
      <div className="max-w-7xl mx-auto px-5 2xl:px-0 w-full">
        <div className="grid grid-cols-2">{id}</div>
        <div>{JSON.stringify(spaceData)}</div>
        <Textarea onChange={(e) => setDescription(e.target.value)}></Textarea>
        <Button label="Make Booking Request" onClick={handleMakeBooking} />
      </div>
    </DefaultLayout>
  );
}
