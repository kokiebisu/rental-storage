import { ListingServiceImpl } from "../../../../App/Service";

export const handler = async (event: any) => {
  const { uid } = event.arguments;
  const service = await ListingServiceImpl.create();
  return await service.removeListingById(uid);
};
