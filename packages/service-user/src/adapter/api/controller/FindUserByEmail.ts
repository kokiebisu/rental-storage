import { UserServiceImpl } from "../../service";

export const handler = async (event: any) => {
  console.log("FIND USER BY EMAIL EVENT: ", event);
  const { email } = event.queryStringParameters;
  const service = await UserServiceImpl.create();
  const userDTO = service.findByEmail(email);
  return {
    statusCode: 200,
    body: JSON.stringify(userDTO),
  };
};
