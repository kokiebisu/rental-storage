import { UserServiceImpl } from "../../../../../application/service/UserService";

export const handler = async (event: any) => {
  const { id } = event.arguments;
  const service = await UserServiceImpl.create();
  return service.findUserById(id);
};
