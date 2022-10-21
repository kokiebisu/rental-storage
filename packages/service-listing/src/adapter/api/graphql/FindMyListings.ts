import { ListingServiceImpl } from "../../../app/service";

export const handler = async (event: any) => {
  const { uid } = event.identity.resolverContext;
  const service = await ListingServiceImpl.create();
  return service.findListingsByUserId(uid);
};
