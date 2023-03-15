import { AppSyncIdentityLambda, AppSyncResolverEvent } from "aws-lambda";
import { FindBookingCommand, FindBookingUseCase } from "../usecase/findBooking";
import { FindSpaceCommand, FindSpaceUseCase } from "../usecase/findSpace";
import { FindSpacesCommand, FindSpacesUseCase } from "../usecase/findSpaces";
import { FindMeCommand, FindMeUseCase } from "../usecase/findMe";
import { FindUserCommand, FindUserUseCase } from "../usecase/findUser";
import {
  GetPresignedURLCommand,
  GetPresignedURLUseCase,
} from "../usecase/getPresignedURL";
import {
  FindBookingsCommand,
  FindBookingsUseCase,
} from "../usecase/findBookings";
import { LoggerClient } from "../../client";

const logger = new LoggerClient();

export const findBooking = async (
  event: AppSyncResolverEvent<{ id: string }, unknown>
) => {
  logger.info(event.arguments, __filename, 22);
  const usecase = new FindBookingUseCase();
  const response = await usecase.execute(
    new FindBookingCommand(event.arguments)
  );
  logger.info(response, __filename, 27);
  return response;
};

export const findBookings = async (
  event: AppSyncResolverEvent<
    { spaceId: string; bookingStatus: "pending" | "approved" },
    unknown
  >
) => {
  logger.info(event.arguments, __filename, 37);
  const usecase = new FindBookingsUseCase();
  const response = await usecase.execute(
    new FindBookingsCommand(event.arguments)
  );
  logger.info(response, __filename, 42);
  return response;
};

export const findSpace = async (
  event: AppSyncResolverEvent<{ id: string }, unknown>
): Promise<ISpace> => {
  logger.info(event.arguments, __filename, 49);
  const usecase = new FindSpaceUseCase();
  const response = await usecase.execute(new FindSpaceCommand(event.arguments));
  logger.info(response, __filename, 52);
  return response;
};

export const findSpaces = async (
  event: AppSyncResolverEvent<
    { filter?: { userId?: string; offset?: number; limit?: number } },
    unknown
  >
) => {
  logger.info(event.arguments, __filename, 59);
  const usecase = new FindSpacesUseCase();

  const response = await usecase.execute(
    new FindSpacesCommand({
      ...(event.arguments.filter?.userId && {
        userId: event.arguments.filter.userId,
      }),
      ...(event.arguments.filter?.offset !== undefined && {
        offset: event.arguments.filter.offset,
      }),
      ...(event.arguments.filter?.limit !== undefined && {
        limit: event.arguments.filter.limit,
      }),
    })
  );
  logger.info(response, __filename, 66);
  return response;
};

export const getPresignedURL = async (
  event: AppSyncResolverEvent<{ filename: string }, unknown>
) => {
  logger.info(event.arguments, __filename, 73);
  const usecase = new GetPresignedURLUseCase();
  const response = await usecase.execute(
    new GetPresignedURLCommand(event.arguments)
  );
  logger.info(response, __filename, 78);
  return response;
};

export const findMe = async (
  event: AppSyncResolverEvent<Record<string, never>, unknown>
) => {
  logger.info(event.arguments, __filename, 85);
  const usecase = new FindMeUseCase();
  const response = await usecase.execute(
    new FindMeCommand({
      id: (event.identity as AppSyncIdentityLambda).resolverContext.uid,
    })
  );
  logger.info(response, __filename, 92);
  return response;
};

export const findUser = async (
  event: AppSyncResolverEvent<{ id: string }, unknown>
) => {
  logger.info(event.arguments, __filename, 99);
  const usecase = new FindUserUseCase();
  const response = await usecase.execute(
    new FindUserCommand({ id: event.arguments.id })
  );
  logger.info(response, __filename, 104);
  return response;
};
