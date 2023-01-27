import { AppSyncResolverEvent } from "aws-lambda";
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
      fee: unknown;
    },
    unknown
  >
) => {
  const usecase = new AddListingUseCase();
  return await usecase.execute(new AddListingCommand(event.arguments));
};

export const removeListingById = async (
  event: AppSyncResolverEvent<{ listingId: string }, unknown>
) => {
  const usecase = new RemoveListingByIdUseCase();
  return await usecase.execute(new RemoveListingByIdCommand(event.arguments));
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
  const usecase = new MakeBookingUseCase();
  return await usecase.execute(new MakeBookingCommand(event.arguments));
};

export const removeUserById = async (
  event: AppSyncResolverEvent<{ userId: string }, unknown>
) => {
  const usecase = new RemoveUserByIdUseCase();
  return await usecase.execute(new RemoveUserByIdCommand(event.arguments));
};
