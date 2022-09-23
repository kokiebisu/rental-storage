import { UserServiceImpl } from "../../../service";

export const handler = async (event: any) => {
  const { id } = event.arguments;
  const service = await UserServiceImpl.create();
  return await service.removeGuestById(id);
};
