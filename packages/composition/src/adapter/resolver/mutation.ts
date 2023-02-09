import { AppSyncIdentityLambda, AppSyncResolverEvent } from "aws-lambda";
import { isCustomError } from "../../helper";
import { AddListingCommand, AddListingUseCase } from "../usecase/addListing";
import { MakeBookingCommand, MakeBookingUseCase } from "../usecase/makeBooking";
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
  try {
    const uid = (event.identity as AppSyncIdentityLambda).resolverContext.uid;
    const input = { ...event.arguments, lenderId: uid };
    const usecase = new AddListingUseCase();
    const result = await usecase.execute(new AddListingCommand(input));
    return result;
  } catch (err: unknown) {
    return isCustomError(err) ? err.serializeError() : err;
  }
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

export const makeBooking = async (
  event: AppSyncResolverEvent<
    {
      userId: string;
      listingId: string;
      items: BookingItem[];
    },
    unknown
  >
) => {
  try {
    const uid = (event.identity as AppSyncIdentityLambda).resolverContext.uid;
    const input = { ...event.arguments, lenderId: uid };
    const usecase = new MakeBookingUseCase();
    return await usecase.execute(new MakeBookingCommand(input));
  } catch (err: unknown) {
    return isCustomError(err) ? err.serializeError() : err;
  }
};

export const removeUserById = async (
  event: AppSyncResolverEvent<{ userId: string }, unknown>
) => {
  try {
    const input = { ...event.arguments };
    const usecase = new RemoveUserByIdUseCase();
    return await usecase.execute(new RemoveUserByIdCommand(input));
  } catch (err: unknown) {
    return isCustomError(err) ? err.serializeError() : err;
  }
};
