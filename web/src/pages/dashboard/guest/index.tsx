import { useQuery } from "@apollo/client";

import { awsLambdaClient } from "@/clients";
import { FIND_MY_BOOKINGS_QUERY } from "@/graphql/queries";
import { DefaultLayout } from "@/layout";
import { Booking } from "@/types/interface";

export default function Dashboard() {
  const { data, loading, error } = useQuery(FIND_MY_BOOKINGS_QUERY, {
    client: awsLambdaClient,
  });
  if (loading) {
    return <div>loading...</div>;
  }
  if (error) {
    return <div>error...</div>;
  }

  console.log("DATA: ", data);
  return (
    <DefaultLayout>
      <div className="w-full h-full">
        <div>
          <div>Pending Bookings</div>
          <div>
            {data.profile.bookings.pending.length > 0
              ? data.profile.bookings.pending.map((booking: Booking) => {
                  return <div key={booking.id}>{JSON.stringify(booking)}</div>;
                })
              : null}
          </div>
        </div>
        <div>
          <div>Approved Bookings</div>
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
