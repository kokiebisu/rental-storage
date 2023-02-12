import { AppSyncIdentityLambda, AppSyncResolverEvent } from "aws-lambda";
import { isCustomError } from "../../helper";
import { FindBookingCommand, FindBookingUseCase } from "../usecase/findBooking";
import { FindListingCommand, FindListingUseCase } from "../usecase/findListing";
import {
  FindListingsCommand,
  FindListingsUseCase,
} from "../usecase/findListings";
import { FindMeCommand, FindMeUseCase } from "../usecase/findMe";
import { FindUserCommand, FindUserUseCase } from "../usecase/findUser";
import {
  GetPresignedURLCommand,
  GetPresignedURLUseCase,
} from "../usecase/getPresignedURL";

export const findBooking = async (
  event: AppSyncResolverEvent<{ id: string }, unknown>
) => {
  const usecase = new FindBookingUseCase();
  return await usecase.execute(new FindBookingCommand(event.arguments));
};

export const findListing = async (
  event: AppSyncResolverEvent<{ id: string }, unknown>
) => {
  const usecase = new FindListingUseCase();
  return await usecase.execute(new FindListingCommand(event.arguments));
};

export const findListings = async (
  event: AppSyncResolverEvent<{ userId: string }, unknown>
) => {
  const usecase = new FindListingsUseCase();
  return await usecase.execute(
    new FindListingsCommand({
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
