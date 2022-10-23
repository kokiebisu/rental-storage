import { ItemService } from "../../../src/app/port";
import { ItemServiceImpl } from "../../../src/app/service";

let service: ItemService;

beforeAll(async () => {
  service = await ItemServiceImpl.create();
});

describe("addItem()", () => {
  it("should work with valid input", () => {});
});
