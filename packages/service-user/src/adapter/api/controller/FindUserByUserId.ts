import { UserServiceImpl } from "../../service";

export const handler = async (event: any) => {
  const { userId } = event.pathParameters;
  const service = await UserServiceImpl.create();
  const user = await service.findById(userId);
  return { statusCode: 200, body: JSON.stringify(user) };
};
