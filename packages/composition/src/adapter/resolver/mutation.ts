import { AppSyncIdentityLambda, AppSyncResolverEvent } from "aws-lambda";
import {
  CreateListingCommand,
  CreateListingUseCase,
} from "../usecase/createListing";
import {
  CreateBookingCommand,
  CreateBookingUseCase,
} from "../usecase/createBooking";
import {
  DeleteListingCommand,
  DeleteListingUseCase,
} from "../usecase/deleteListing";
import { DeleteUserCommand, DeleteUserUseCase } from "../usecase/deleteUser";

export const createListing = async (
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
  const usecase = new CreateListingUseCase();
  return await usecase.execute(new CreateListingCommand(input));
};

export const deleteListing = async (
  event: AppSyncResolverEvent<{ id: string }, unknown>
) => {
  const usecase = new DeleteListingUseCase();
  return await usecase.execute(new DeleteListingCommand(event.arguments));
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

export const deleteUser = async (
  event: AppSyncResolverEvent<{ id: string }, unknown>
) => {
  const input = { ...event.arguments };
  const usecase = new DeleteUserUseCase();
  return await usecase.execute(new DeleteUserCommand(input));
};
