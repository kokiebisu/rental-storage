import { AppSyncIdentityLambda, AppSyncResolverEvent } from "aws-lambda";
import { isCustomError } from "../../helper";
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

export const findBooking = async (
  event: AppSyncResolverEvent<{ id: string }, unknown>
) => {
  const usecase = new FindBookingUseCase();
  return await usecase.execute(new FindBookingCommand(event.arguments));
};

export const findBookings = async (
  event: AppSyncResolverEvent<
    { spaceId: string; bookingStatus: "PENDING" | "APPROVED" },
    unknown
  >
) => {
  const usecase = new FindBookingsUseCase();
  return await usecase.execute(new FindBookingsCommand(event.arguments));
};

export const findSpace = async (
  event: AppSyncResolverEvent<{ id: string }, unknown>
): Promise<ISpace> => {
  const usecase = new FindSpaceUseCase();
  return await usecase.execute(new FindSpaceCommand(event.arguments));
};

export const findSpaces = async (
  event: AppSyncResolverEvent<{ userId: string }, unknown>
) => {
  const usecase = new FindSpacesUseCase();
  return await usecase.execute(
    new FindSpacesCommand({
      userId: event.arguments.userId,
    })
  );
};

export const getPresignedURL = async (
  event: AppSyncResolverEvent<{ filename: string }, unknown>
) => {
  try {
    const usecase = new GetPresignedURLUseCase();
    return await usecase.execute(new GetPresignedURLCommand(event.arguments));
  } catch (err: unknown) {
    return isCustomError(err) ? err.serializeError() : err;
  }
};

export const findMe = async (
  event: AppSyncResolverEvent<Record<string, never>, unknown>
) => {
  const usecase = new FindMeUseCase();
  return await usecase.execute(
    new FindMeCommand({
      id: (event.identity as AppSyncIdentityLambda).resolverContext.uid,
    })
  );
};

export const findUser = async (
  event: AppSyncResolverEvent<{ id: string }, unknown>
) => {
  const usecase = new FindUserUseCase();
  const response = await usecase.execute(
    new FindUserCommand({ id: event.arguments.id })
  );
  return response;
};
