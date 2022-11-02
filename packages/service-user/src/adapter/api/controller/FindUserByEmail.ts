import { UserServiceImpl } from "../../service";

export const handler = async (event: any) => {
  console.log("FIND USER BY EMAIL EVENT: ", event);
  const { uid } = event.arguments;
  const service = await UserServiceImpl.create();
  return service.findById(uid);
};
