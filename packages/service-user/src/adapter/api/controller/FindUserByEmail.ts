import { UserServiceImpl } from "../../service";

export const handler = async (event: any) => {
  console.log("FIND USER BY EMAIL EVENT: ", event);
  const { email } = event.queryStringParameters;
  const service = await UserServiceImpl.create();
<<<<<<< HEAD
  return service.findByEmail(email);
=======
  const userDTO = service.findByEmail(email);
  return {
    statusCode: 200,
    body: JSON.stringify(userDTO),
  };
>>>>>>> 313874a240b3c3f9b8e03735c93949ca1277492f
};
