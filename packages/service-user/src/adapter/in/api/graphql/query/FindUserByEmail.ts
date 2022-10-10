import { UserServiceImpl } from "../../../../../application/service/UserService";

export const handler = async (event: any) => {
  console.log("LAMBDA EVENT: ", event);
  const { email } = event.arguments;
  const service = await UserServiceImpl.create();
  return service.findUserByEmail(email);
};
