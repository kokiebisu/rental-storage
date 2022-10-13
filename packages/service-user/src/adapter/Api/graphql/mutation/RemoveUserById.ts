import { UserServiceImpl } from "../../../../application/Service/UserService";

export const handler = async (event: any) => {
  const { uid } = event.arguments;
  const service = await UserServiceImpl.create();
  return await service.removeById(uid);
};
