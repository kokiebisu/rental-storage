import { useQuery } from "@apollo/client";
import { useRouter } from "next/router";

import { apiKeyClient } from "@/clients";
import { Grid } from "@/components";
import { DefaultLayout } from "@/layout";
import { FIND_SPACE_QUERY } from "@/queries";

export default function SpaceDetailsPage() {
  const router = useRouter();
  const { id } = router.query;
  const { data, error, loading } = useQuery(FIND_SPACE_QUERY, {
    variables: { id },
    client: apiKeyClient,
  });

  if (loading) {
    return <div>loading</div>;
  }

  if (error) {
    return <div>error</div>;
  }
  return (
    <DefaultLayout>
      <div className="my-6">
        <Grid />
      </div>
      <div className="max-w-7xl mx-auto px-5 2xl:px-0 w-full">
        <div className="grid grid-cols-2">{id}</div>
        <div>{JSON.stringify(data)}</div>
      </div>
    </DefaultLayout>
  );
}
