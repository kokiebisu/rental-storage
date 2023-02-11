import { AppSyncIdentityLambda, AppSyncResolverEvent } from "aws-lambda";
import { isCustomError } from "../../helper";
import { AddListingCommand, AddListingUseCase } from "../usecase/addListing";
import {
  CreateBookingCommand,
  CreateBookingUseCase,
} from "../usecase/createBooking";
import {
  RemoveListingByIdCommand,
  RemoveListingByIdUseCase,
} from "../usecase/removeListingById";
import {
  RemoveUserByIdCommand,
  RemoveUserByIdUseCase,
} from "../usecase/removeUserById";

export const addListing = async (
  event: AppSyncResolverEvent<
    {
      streetAddress: string;
      latitude: number;
      longitude: number;
      imageUrls: string[];
      title: string;
      feeAmount: number;
      feeCurrency: string;
      feeType: string;
    },
    unknown
  >
) => {
  const uid = (event.identity as AppSyncIdentityLambda).resolverContext.uid;
  const input = { ...event.arguments, lenderId: uid };
  const usecase = new AddListingUseCase();
  return await usecase.execute(new AddListingCommand(input));
};

export const removeListingById = async (
  event: AppSyncResolverEvent<{ listingId: string }, unknown>
) => {
  try {
    const input = event.arguments;
    const usecase = new RemoveListingByIdUseCase();
    return await usecase.execute(new RemoveListingByIdCommand(input));
  } catch (err: unknown) {
    return isCustomError(err) ? err.serializeError() : err;
  }
};

export const createBooking = async (
  event: AppSyncResolverEvent<
    {
      listingId: string;
      items: BookingItem[];
    },
    unknown
  >
) => {
  const uid = (event.identity as AppSyncIdentityLambda).resolverContext.uid;
  const input = { ...event.arguments, userId: uid };
  const usecase = new CreateBookingUseCase();
  return await usecase.execute(new CreateBookingCommand(input));
};

export const removeUserById = async (
  event: AppSyncResolverEvent<{ userId: string }, unknown>
) => {
  const input = { ...event.arguments };
  const usecase = new RemoveUserByIdUseCase();
  return await usecase.execute(new RemoveUserByIdCommand(input));
};
