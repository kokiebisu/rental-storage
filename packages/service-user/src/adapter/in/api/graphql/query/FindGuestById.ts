import { UserServiceImpl } from "../../../../../application/service/UserService";

export const handler = async (event: any) => {
  console.log("DB_HOST: ", process.env.DB_HOST);
  console.log("DB_USERNAME: ", process.env.DB_USERNAME);
  console.log("DB_NAME: ", process.env.DB_NAME);
  console.log("DB_PASSWORD: ", process.env.DB_PASSWORD);
  console.log("SERVICE_API_ENDPOINT: ", process.env.SERVICE_API_ENDPOINT);
  const { id } = event.arguments;
  const service = await UserServiceImpl.create();
  return service.findGuestById(id);
};
