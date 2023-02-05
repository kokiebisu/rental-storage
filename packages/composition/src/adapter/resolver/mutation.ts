import { AppSyncResolverEvent } from "aws-lambda";
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
      lenderId: string;
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
    const usecase = new AddListingUseCase();
    return await usecase.execute(new AddListingCommand(event.arguments));
  } catch (err: unknown) {
    return isCustomError(err) ? err.serializeError() : err;
  }
};

export const removeListingById = async (
  event: AppSyncResolverEvent<{ listingId: string }, unknown>
) => {
  try {
    const usecase = new RemoveListingByIdUseCase();
    return await usecase.execute(new RemoveListingByIdCommand(event.arguments));
  } catch (err: unknown) {
    return isCustomError(err) ? err.serializeError() : err;
  }
};

export const makeBooking = async (
  event: AppSyncResolverEvent<
    {
      userId: string;
      amount: number;
      currency: string;
      listingId: string;
      items: unknown;
    },
    unknown
  >
) => {
  try {
    const usecase = new MakeBookingUseCase();
    return await usecase.execute(new MakeBookingCommand(event.arguments));
  } catch (err: unknown) {
    return isCustomError(err) ? err.serializeError() : err;
  }
};

export const removeUserById = async (
  event: AppSyncResolverEvent<{ userId: string }, unknown>
) => {
  try {
    const usecase = new RemoveUserByIdUseCase();
    return await usecase.execute(new RemoveUserByIdCommand(event.arguments));
  } catch (err: unknown) {
    return isCustomError(err) ? err.serializeError() : err;
  }
};
