import { UserServiceImpl } from "../../service";

export const handler = async (event: any) => {
  const { emailAddress } = event.queryStringParameters;
  const service = await UserServiceImpl.create();
  const userDTO = await service.findByEmail(emailAddress);
  return {
    statusCode: 200,
    body: JSON.stringify(userDTO),
  };
};
