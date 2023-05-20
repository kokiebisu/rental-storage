import { AppSyncIdentityLambda, AppSyncResolverEvent } from "aws-lambda";
import { CreateSpaceCommand, CreateSpaceUseCase } from "../usecase/createSpace";
import {
  CreateBookingCommand,
  CreateBookingUseCase,
} from "../usecase/createBooking";
import { DeleteSpaceCommand, DeleteSpaceUseCase } from "../usecase/deleteSpace";
import { DeleteUserCommand, DeleteUserUseCase } from "../usecase/deleteUser";
import { LoggerClient } from "../../client";
import {
  AcceptBookingCommand,
  AcceptBookingUseCase,
} from "../usecase/acceptBooking";

const logger = new LoggerClient();

export const createSpace = async (
  event: AppSyncResolverEvent<
    {
      title: string;
      description: string;
      imageUrls: string[];
      location: ILocation;
    },
    unknown
  >
) => {
  logger.info(event.arguments, __filename, 28);
  const uid = (event.identity as AppSyncIdentityLambda).resolverContext.uid;
  const input = { ...event.arguments, lenderId: uid };
  const usecase = new CreateSpaceUseCase();
  const response = await usecase.execute(new CreateSpaceCommand(input));
  logger.info(response, __filename, 33);
  return response;
};

export const deleteSpace = async (
  event: AppSyncResolverEvent<{ id: string }, unknown>
) => {
  logger.info(event.arguments, __filename, 36);
  const usecase = new DeleteSpaceUseCase();
  const response = await usecase.execute(
    new DeleteSpaceCommand(event.arguments)
  );
  logger.info(response, __filename, 41);
  return response;
};

export const createBooking = async (
  event: AppSyncResolverEvent<
    {
      spaceId: string;
      imageUrls: string[];
      description: string;
      startDate: string;
      endDate: string;
    },
    unknown
  >
) => {
  logger.info(event.arguments, __filename, 57);
  const uid = (event.identity as AppSyncIdentityLambda).resolverContext.uid;
  const input = { ...event.arguments, userId: uid };
  const usecase = new CreateBookingUseCase();
  const response = await usecase.execute(new CreateBookingCommand(input));
  logger.info(response, __filename, 62);
  return response;
};

export const acceptBooking = async (
  event: AppSyncResolverEvent<
    {
      id: string;
    },
    unknown
  >
) => {
  logger.info(event.arguments, __filename, 74);
  const input = { ...event.arguments };
  const usecase = new AcceptBookingUseCase();
  const response = await usecase.execute(new AcceptBookingCommand(input));
  logger.info(response, __filename, 62);
  return response;
};

export const deleteUser = async (
  event: AppSyncResolverEvent<{ id: string }, unknown>
) => {
  logger.info(event.arguments, __filename, 69);
  const input = { ...event.arguments };
  const usecase = new DeleteUserUseCase();
  const response = await usecase.execute(new DeleteUserCommand(input));
  logger.info(response, __filename, 73);
  return response;
};
