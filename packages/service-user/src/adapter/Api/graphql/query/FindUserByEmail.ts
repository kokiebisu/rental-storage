import { UserServiceImpl } from "../../../../application/Service/UserService";

export const handler = async (event: any) => {
  const { email } = event.arguments;
  const service = await UserServiceImpl.create();
  return service.findByEmail(email);
};
