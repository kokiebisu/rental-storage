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

export const findListingById = async (event: any) => {
  const usecase = new FindListingByIdUseCase();
  return await usecase.execute(new FindListingByIdCommand(event.arguments));
};

export const findListingsWithinLatLng = async (event: any) => {
  const usecase = new FindListingsWithinLatLngUseCase();
  return await usecase.execute(
    new FindListingsWithinLatLngCommand(event.arguments)
  );
};

export const findMyListings = async (event: any) => {
  const usecase = new FindMyListingsUseCase();
  return await usecase.execute(
    new FindMyListingsCommand(event.identity.resolverContext)
  );
};

export const getPresignedURL = async (event: any) => {
  const usecase = new GetPresignedURLUseCase();
  return await usecase.execute(new GetPresignedURLCommand(event.arguments));
};

export const findAllCreatedBookings = async (event: any) => {
  const usecase = new FindAllCreatedBookingsUseCase();
  return await usecase.execute(
    new FindAllCreatedBookingsCommand(event.identity.resolverContext)
  );
};

export const findMe = async (event: any) => {
  const usecase = new FindMeUseCase();
  return await usecase.execute(
    new FindMeCommand({ userId: event.identity.resolverContext.uid })
  );
};

export const findUserByEmail = async (event: any) => {
  const usecase = new FindUserByEmailUseCase();
  return await usecase.execute(new FindUserByEmailCommand(event.arguments));
};

export const findUserById = async (event: any) => {
  const usecase = new FindUserByIdUseCase();
  return await usecase.execute(
    new FindUserByIdCommand({ userId: event.identity.resolverContext.uid })
  );
};
