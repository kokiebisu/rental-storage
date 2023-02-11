import { AppSyncIdentityLambda, AppSyncResolverEvent } from "aws-lambda";
import { isCustomError } from "../../helper";

import {
  FindAllCreatedBookingsCommand,
  FindAllCreatedBookingsUseCase,
} from "../usecase/findAllCreatedBookings";
import {
  FindBookingByIdCommand,
  FindBookingByIdUseCase,
} from "../usecase/findBookingById";
import {
  FindListingByIdCommand,
  FindListingByIdUseCase,
} from "../usecase/findListingById";
import {
  FindListingsByUserIdCommand,
  FindListingsByUserIdUseCase,
} from "../usecase/findListingsByUserId";
import {
  FindListingsWithinLatLngCommand,
  FindListingsWithinLatLngUseCase,
} from "../usecase/findListingsWithinLatLng";
import { FindMeCommand, FindMeUseCase } from "../usecase/findMe";
import {
  FindUserByEmailCommand,
  FindUserByEmailUseCase,
} from "../usecase/findUserByEmail";
import {
  FindUserByIdCommand,
  FindUserByIdUseCase,
} from "../usecase/findUserById";
import {
  GetPresignedURLCommand,
  GetPresignedURLUseCase,
} from "../usecase/getPresignedURL";

export const findBookingById = async (
  event: AppSyncResolverEvent<{ bookingId: string }, unknown>
) => {
  const usecase = new FindBookingByIdUseCase();
  return await usecase.execute(new FindBookingByIdCommand(event.arguments));
};

export const findListingById = async (
  event: AppSyncResolverEvent<{ listingId: string }, unknown>
) => {
  const usecase = new FindListingByIdUseCase();
  return await usecase.execute(new FindListingByIdCommand(event.arguments));
};

export const findListingsWithinLatLng = async (
  event: AppSyncResolverEvent<
    { latitude: number; longitude: number; range: number },
    unknown
  >
) => {
  try {
    const usecase = new FindListingsWithinLatLngUseCase();
    return await usecase.execute(
      new FindListingsWithinLatLngCommand(event.arguments)
    );
  } catch (err: unknown) {
    return isCustomError(err) ? err.serializeError() : err;
  }
};

export const findMyListings = async (
  event: AppSyncResolverEvent<Record<string, never>, unknown>
) => {
  try {
    const usecase = new FindListingsByUserIdUseCase();
    return await usecase.execute(
      new FindListingsByUserIdCommand({
        userId: (event.identity as AppSyncIdentityLambda).resolverContext,
      })
    );
  } catch (err: unknown) {
    return isCustomError(err) ? err.serializeError() : err;
  }
};

export const findListingsByUserId = async (
  event: AppSyncResolverEvent<{ userId: string }, unknown>
) => {
  const usecase = new FindListingsByUserIdUseCase();
  return await usecase.execute(
    new FindListingsByUserIdCommand({
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

export const findAllCreatedBookings = async (
  event: AppSyncResolverEvent<Record<string, never>, unknown>
) => {
  try {
    const usecase = new FindAllCreatedBookingsUseCase();
    return await usecase.execute(
      new FindAllCreatedBookingsCommand(
        (event.identity as AppSyncIdentityLambda).resolverContext
      )
    );
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
      userId: (event.identity as AppSyncIdentityLambda).resolverContext.uid,
    })
  );
};

export const findUserByEmail = async (
  event: AppSyncResolverEvent<{ emailAddress: string }, unknown>
) => {
  const usecase = new FindUserByEmailUseCase();
  return await usecase.execute(new FindUserByEmailCommand(event.arguments));
};

export const findUserById = async (
  event: AppSyncResolverEvent<{ userId: string }, unknown>
) => {
  try {
    const usecase = new FindUserByIdUseCase();
    const response = await usecase.execute(
      new FindUserByIdCommand({ userId: event.arguments.userId })
    );
    return response;
  } catch (err: unknown) {
    return isCustomError(err) ? err.serializeError() : err;
  }
};
