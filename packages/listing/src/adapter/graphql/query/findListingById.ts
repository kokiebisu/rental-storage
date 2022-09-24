import { ListingServiceImpl } from "../../../service";

export const handler = async (event: any) => {
  const { id } = event.arguments;
  const service = await ListingServiceImpl.create();
  return service.findListingById(id);
};
