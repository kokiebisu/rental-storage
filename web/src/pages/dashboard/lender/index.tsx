import { awsLambdaClient } from "@/clients";
import { Button } from "@/components";
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
    return <div>loading...</div>;
  }
  if (error) {
    return <div>error...</div>;
  }

  return (
    <DefaultLayout>
      <div>
        <div className="mx-auto mt-10 flex justify-center">
          <Link href="/dashboard/lender/create">
            <Button onClick={() => {}} label="Add new space" />
          </Link>
        </div>
        <div>
          {data.spaces.length > 0
            ? data.spaces.map((space: Space) => {
                return <div key={space.id}>{JSON.stringify(space)}</div>;
              })
            : null}
        </div>
      </div>
    </DefaultLayout>
  );
}
