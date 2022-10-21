import { ListingServiceImpl } from "../../../apps/service";

export const handler = async (event: any) => {
  const { uid } = event.arguments;
  const service = await ListingServiceImpl.create();
  return service.findListingById(uid);
};
