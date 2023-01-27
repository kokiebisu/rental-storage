import { AppSyncIdentityLambda, AppSyncResolverEvent } from "aws-lambda";

import {
  FindAllCreatedBookingsCommand,
  FindAllCreatedBookingsUseCase,
} from "../usecase/findAllCreatedBookings";
import {
  FindListingByIdCommand,
  FindListingByIdUseCase,
} from "../usecase/findListingById";
import {
  FindListingsWithinLatLngCommand,
  FindListingsWithinLatLngUseCase,
} from "../usecase/findListingsWithinLatLng";
import { FindMeCommand, FindMeUseCase } from "../usecase/findMe";
import {
  FindMyListingsCommand,
  FindMyListingsUseCase,
} from "../usecase/findMyListings";
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
  const usecase = new FindListingsWithinLatLngUseCase();
  return await usecase.execute(
    new FindListingsWithinLatLngCommand(event.arguments)
  );
};

export const findMyListings = async (
  event: AppSyncResolverEvent<{ uid: string }, unknown>
) => {
  const usecase = new FindMyListingsUseCase();
  return await usecase.execute(
    new FindMyListingsCommand(
      (event.identity as AppSyncIdentityLambda).resolverContext
    )
  );
};

export const getPresignedURL = async (
  event: AppSyncResolverEvent<{ filename: string }, unknown>
) => {
  const usecase = new GetPresignedURLUseCase();
  return await usecase.execute(new GetPresignedURLCommand(event.arguments));
};

export const findAllCreatedBookings = async (
  event: AppSyncResolverEvent<Record<string, never>, unknown>
) => {
  const usecase = new FindAllCreatedBookingsUseCase();
  return await usecase.execute(
    new FindAllCreatedBookingsCommand(
      (event.identity as AppSyncIdentityLambda).resolverContext
    )
  );
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
  event: AppSyncResolverEvent<{ email: string }, unknown>
) => {
  const usecase = new FindUserByEmailUseCase();
  return await usecase.execute(new FindUserByEmailCommand(event.arguments));
};

export const findUserById = async (
  event: AppSyncResolverEvent<{ uid: string }, unknown>
) => {
  const usecase = new FindUserByIdUseCase();
  const response = await usecase.execute(
    new FindUserByIdCommand({ userId: event.arguments.uid })
  );
  return response;
};
