import { UserServiceImpl } from "../../service";

export const handler = async (event: any) => {
  console.log("CREATE USER EVENT: ", event);
  const { firstName, lastName, emailAddress, password } =
    event.queryStringParameters;
  const service = await UserServiceImpl.create();

  return service.createUser({ firstName, lastName, emailAddress, password });
};
