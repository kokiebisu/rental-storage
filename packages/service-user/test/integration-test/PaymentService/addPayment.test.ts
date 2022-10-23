import { ListingService } from "../../../../service-listing/src/app/port";
import { ListingServiceImpl } from "../../../../service-listing/src/app/service";

let service: ListingService;

beforeAll(async () => {
  service = await ListingServiceImpl.create();
});

describe("addPayment()", () => {
  it("should work with valid input", () => {});
});
