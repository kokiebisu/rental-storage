import { UserServiceImpl } from "../../../app/service/UserService";

export const handler = async (event: any) => {
  const { uid } = event.identity.resolverContext;
  const service = await UserServiceImpl.create();
  return service.findById(uid);
};
