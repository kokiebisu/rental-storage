import { UserServiceImpl } from "../../service";

export const handler = async (event: any) => {
  console.log("CREATE USER EVENT: ", event);
  const { firstName, lastName, emailAddress, password } = JSON.parse(
    event.body
  );
  const service = await UserServiceImpl.create();

  const userDTO = await service.createUser({
    firstName,
    lastName,
    emailAddress,
    password,
  });

  return {
    statusCode: 200,
    body: JSON.stringify(userDTO),
  };
};
