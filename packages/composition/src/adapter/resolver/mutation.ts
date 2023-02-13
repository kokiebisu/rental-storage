import { AppSyncIdentityLambda, AppSyncResolverEvent } from "aws-lambda";
import { CreateSpaceCommand, CreateSpaceUseCase } from "../usecase/createSpace";
import {
  CreateBookingCommand,
  CreateBookingUseCase,
} from "../usecase/createBooking";
import { DeleteSpaceCommand, DeleteSpaceUseCase } from "../usecase/deleteSpace";
import { DeleteUserCommand, DeleteUserUseCase } from "../usecase/deleteUser";

export const createSpace = async (
  event: AppSyncResolverEvent<
    {
      streetAddress: string;
      latitude: number;
      longitude: number;
      imageUrls: string[];
      title: string;
      description: string;
    },
    unknown
  >
) => {
  const uid = (event.identity as AppSyncIdentityLambda).resolverContext.uid;
  const input = { ...event.arguments, lenderId: uid };
  const usecase = new CreateSpaceUseCase();
  return await usecase.execute(new CreateSpaceCommand(input));
};

export const deleteSpace = async (
  event: AppSyncResolverEvent<{ id: string }, unknown>
) => {
  const usecase = new DeleteSpaceUseCase();
  return await usecase.execute(new DeleteSpaceCommand(event.arguments));
};

export const createBooking = async (
  event: AppSyncResolverEvent<
    {
      spaceId: string;
      startDate: string;
      endDate: string;
    },
    unknown
  >
) => {
  const uid = (event.identity as AppSyncIdentityLambda).resolverContext.uid;
  const input = { ...event.arguments, userId: uid };
  const usecase = new CreateBookingUseCase();
  return await usecase.execute(new CreateBookingCommand(input));
};

export const deleteUser = async (
  event: AppSyncResolverEvent<{ id: string }, unknown>
) => {
  const input = { ...event.arguments };
  const usecase = new DeleteUserUseCase();
  return await usecase.execute(new DeleteUserCommand(input));
};
