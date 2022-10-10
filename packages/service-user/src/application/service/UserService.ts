import axios from "axios";

import { EmailAddress, Payment, User } from "../../domain/model";
import { CreateUserInput, UserRepository, UserService } from "../port";

import { LoggerUtil } from "../../utils";
import { UserInterface } from "../../types";
import { UserRepositoryImpl } from "../../adapter/out/db";
import { UserMapper } from "../../adapter/in/mapper";

export interface AddPaymentInput {
  userId: number;
  emailAddress: string;
  firstName: string;
  lastName: string;
}

export class UserServiceImpl implements UserService {
  private _userRepository: UserRepository;

  private _logger: LoggerUtil;

  private constructor(userRepository: UserRepository) {
    this._userRepository = userRepository;
    this._logger = new LoggerUtil("UserServiceImpl");
  }

  public static async create() {
    const userRepository = await UserRepositoryImpl.create();
    await userRepository.setup();
    return new UserServiceImpl(userRepository);
  }

  public async createUser(data: CreateUserInput): Promise<UserInterface> {
    this._logger.info(data, "createUser()");
    try {
      let user = new User({
        firstName: data.firstName,
        lastName: data.lastName,
        emailAddress: new EmailAddress(data?.emailAddress),
        password: data.password,
      });

      const savedUser = await this._userRepository.save(user);

      const userDTO = UserMapper.toDTOFromEntity(savedUser);
      if (!userDTO.id) {
        throw new Error(`id property missing from userDTO`);
      }
      await this.addPayment({
        userId: userDTO.id,
        emailAddress: userDTO.emailAddress,
        firstName: userDTO.firstName,
        lastName: userDTO.lastName,
      });
      return userDTO;
    } catch (err) {
      this._logger.error(err, "createUser()");
      throw err;
    }
  }

  public async addPayment(data: AddPaymentInput) {
    this._logger.info(data, "addPayment()");
    // get stripe customer id for the user
    try {
      const response: { data: { providerId: string; providerType: string } } =
        await axios.post(
          `${process.env.SERVICE_API_ENDPOINT}/payments/customer`,
          {
            emailAddress: data.emailAddress,
            firstName: data.firstName,
            lastName: data.lastName,
          }
        );
      console.log("RESPONSE: ", response);
      const payment = new Payment({
        userId: data.userId,
        providerType: response.data.providerType,
        providerId: response.data.providerId,
      });
      this._userRepository.savePayment(payment);
    } catch (err) {
      console.error(err);
    }
  }

  public async removeUserById(id: number): Promise<boolean> {
    this._logger.info({ id }, "removeUserById()");
    const userExists = await this.findUserById(id);
    if (!userExists) {
      throw new Error(`User with email ${id} doesn't exist`);
    }
    try {
      await this._userRepository.delete(id);
      return true;
    } catch (err) {
      this._logger.error(err, "removeUserById");
      return false;
    }
  }

  public async findUserById(id: number): Promise<UserInterface | null> {
    this._logger.info({ id }, "findUserById()");
    try {
      const user = await this._userRepository.findOneById(id);
      if (!user) {
        throw new Error(`User with id ${id} doesn't exist in db`);
      }
      return UserMapper.toDTOFromEntity(user);
    } catch (err) {
      this._logger.error(err, "findUserById()");
      return null;
    }
  }

  public async findUserByEmail(
    emailAddress: string
  ): Promise<UserInterface | null> {
    this._logger.info({ emailAddress }, "findUserByEmail()");
    try {
      const user = await this._userRepository.findOneByEmail(emailAddress);
      if (!user) {
        throw new Error(`User with email ${emailAddress} doesn't exist in db`);
      }
      return UserMapper.toDTOFromEntity(user);
    } catch (err) {
      this._logger.error(err, "findUserByEmail()");
      return null;
    }
  }
  // MOVE TO items-service
  // public async removeById(id: number): Promise<boolean> {
  //   this._logger.info({ id }, "removeById()");
  //   const userExists = await this.findById(id);
  //   if (!userExists) {
  //     throw new Error(`User with id ${id} doesn't exist`);
  //   }
  //   try {
  //     await this._userRepository.delete(id);
  //     return true;
  //   } catch (err) {
  //     this._logger.error(err, "removeById()");
  //     return false;
  //   }
  // }

  // MOVE TO items-service
  // public async updateStoredItems(booking: BookingInterface): Promise<boolean> {
  //   this._logger.info({ booking }, "updatedStoredItems()");
  //   try {
  //     await this._userRepository.updateStoringItem(
  //       booking.userId,
  //       booking.items
  //     );
  //     return true;
  //   } catch (err) {
  //     this._logger.error(err, "updateStoringItem");
  //     return false;
  //   }
  // }
}
