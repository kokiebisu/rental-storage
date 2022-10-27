import { ListingService } from "../../src/app/port";
import { ListingServiceImpl } from "../../src/app/service";

let service: ListingService;

beforeAll(async () => {
  service = await ListingServiceImpl.create();
});

describe("findListingById()", () => {
  it("should work when input is valid", () => {});
});
