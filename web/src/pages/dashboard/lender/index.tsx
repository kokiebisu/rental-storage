import { awsLambdaClient } from "@/clients";
import { Button } from "@/components";
import { Spinner } from "@/components/spinner";
import { AuthContext } from "@/context/auth";
import { FIND_SPACES_BY_LENDER_QUERY } from "@/graphql/queries/aws-lambda/space.graphql";
import { DefaultLayout } from "@/layout";
import { Space } from "@/types/interface";
import { useQuery } from "@apollo/client";
import Link from "next/link";
import { useContext } from "react";

export default function Dashboard() {
  const { user } = useContext(AuthContext);
  const { data, loading, error } = useQuery(FIND_SPACES_BY_LENDER_QUERY, {
    client: awsLambdaClient,
    variables: {
      filter: {
        userId: user ? user.id : "",
      },
    },
  });
  if (loading) {
    return (
      <div className="absolute top-1/2 left-1/2">
        <Spinner />
      </div>
    );
  }
  if (error) {
    return <div>error...</div>;
  }

  return (
    <DefaultLayout>
      <div className="max-w-7xl mx-auto px-5 2xl:px-0 w-full">
        <div className="h-[256px]">
          <div>
            <h3 className="font-bold text-2xl">Pending Bookings</h3>
          </div>
        </div>
        <div className="h-[256px]">
          <div className="font-bold text-2xl">Approved Bookings</div>
        </div>
        <div className="mx-auto mt-10 flex justify-center">
          <Link href="/dashboard/lender/create">
            <Button onClick={() => {}}>Add new space</Button>
          </Link>
        </div>
        <div>
          {data.spaces.length > 0
            ? data.spaces.map((space: Space, index: number) => {
                return <div key={index}>{JSON.stringify(space)}</div>;
              })
            : null}
        </div>
      </div>
    </DefaultLayout>
  );
}
