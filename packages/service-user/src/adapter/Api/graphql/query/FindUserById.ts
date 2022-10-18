import { UserServiceImpl } from "../../../../Application/Service/UserService";

export const handler = async (event: any) => {
  const { uid } = event.arguments;
  const service = await UserServiceImpl.create();
  return service.findById(uid);
};
