import { UserServiceImpl } from "../../../../application/Service/UserService";

export const handler = async (event: any) => {
  console.log("LAMBDA EVENT: ", event);
  const { email } = event.arguments;
  const service = await UserServiceImpl.create();
  return service.findByEmail(email);
};
