import { ListingService } from "../../src/port/service";
import { ListingServiceImpl } from "../../src/adapter/service";

let service: ListingService;

beforeAll(async () => {
  service = await ListingServiceImpl.create();
});

describe("findListingById()", () => {
  it("should work when input is valid", () => {});
});
