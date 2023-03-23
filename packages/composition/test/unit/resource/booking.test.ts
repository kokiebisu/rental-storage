import { BookingResourceURLBuilder } from "../../../src/resource";

describe("BookingResourceURLBuilder", () => {
  const builder = new BookingResourceURLBuilder();
  it("finds a booking by id", () => {
    const id = "123";
    const result = builder.findBooking(id);
    expect(result).toEqual(`${builder.baseURL}/bookings/${id}`);
  });

  it("finds bookings by spaceId", () => {
    const spaceId = "123";
    const result = builder.findAllBookings({ spaceId });
    expect(result).toEqual(`${builder.baseURL}/bookings?spaceId=${spaceId}`);
  });

  it("finds pending bookings by spaceId", () => {
    const spaceId = "123";
    const result = builder.findPendingBookings({
      spaceId,
    });
    expect(result).toEqual(
      `${builder.baseURL}/bookings?spaceId=${spaceId}&bookingStatus=pending`
    );
  });

  it("finds pending bookings by userId", () => {
    const userId = "123";
    const result = builder.findPendingBookings({
      userId,
    });
    expect(result).toEqual(
      `${builder.baseURL}/bookings?userId=${userId}&bookingStatus=pending`
    );
  });

  it("finds approved bookings by spaceId", () => {
    const spaceId = "123";
    const result = builder.findApprovedBookings({
      spaceId,
    });
    expect(result).toEqual(
      `${builder.baseURL}/bookings?spaceId=${spaceId}&bookingStatus=approved`
    );
  });

  it("acceps a booking by id", () => {
    const bookingId = "123";
    const result = builder.acceptBooking(bookingId);
    expect(result).toEqual(`${builder.baseURL}/bookings/${bookingId}/accept`);
  });
});
