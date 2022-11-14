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

export const addListing = async (event: any) => {
  const usecase = new AddListingUseCase();
  return await usecase.execute(new AddListingCommand(event.arguments));
};

export const removeListingById = async (event: any) => {
  const usecase = new RemoveListingByIdUseCase();
  return await usecase.execute(new RemoveListingByIdCommand(event.arguments));
};

export const makeBooking = async (event: any) => {
  const usecase = new MakeBookingUseCase();
  return await usecase.execute(new MakeBookingCommand(event.arguments));
};

export const removeUserById = async (event: any) => {
  const usecase = new RemoveUserByIdUseCase();
  return await usecase.execute(new RemoveUserByIdCommand(event.arguments));
};
