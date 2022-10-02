import { UserServiceImpl } from "../../../../../application/service/UserService";

export const handler = async (event: any) => {
  const { firstName, lastName, emailAddress, password } = event.arguments;
  const service = await UserServiceImpl.create();
  return await service.registerGuest({
    firstName,
    lastName,
    emailAddress,
    password,
  });
};
