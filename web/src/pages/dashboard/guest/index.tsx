import { useQuery } from "@apollo/client";

import { awsLambdaClient } from "@/clients";
import { FIND_MY_BOOKINGS_QUERY } from "@/graphql/queries";
import { DefaultLayout } from "@/layout";
import { Booking } from "@/types/interface";
import { Spinner } from "@/components/spinner";

export default function Dashboard() {
  const { data, loading, error } = useQuery(FIND_MY_BOOKINGS_QUERY, {
    client: awsLambdaClient,
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
          <div>
            {data.profile.bookings.pending.length > 0
              ? data.profile.bookings.pending.map((booking: Booking) => {
                  return <div key={booking.id}>{JSON.stringify(booking)}</div>;
                })
              : null}
          </div>
        </div>
        <div className="h-[256px]">
          <div className="font-bold text-2xl">Approved Bookings</div>
          <div>
            {data.profile.bookings.approved.length > 0
              ? data.profile.bookings.approved.map((booking: Booking) => {
                  return <div key={booking.id}>{JSON.stringify(booking)}</div>;
                })
              : null}
          </div>
        </div>
      </div>
    </DefaultLayout>
  );
}
