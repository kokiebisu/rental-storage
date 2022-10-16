import { UserServiceImpl } from "../../../../application/Service/UserService";

export const handler = async (event: any) => {
  console.log("FINDME EVENT: ", event);
  const { uid } = event.identity.resolverContext;
  const service = await UserServiceImpl.create();
  return service.findById(uid);
};
