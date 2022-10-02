import { EmailAddress, User } from "../../domain/model";
import { CreateUserInput, UserRepository, UserService } from "../port";

import { LoggerUtil } from "../../utils";
import { BookingInterface, UserInterface } from "../../types";
import { UserRepositoryImpl } from "../../adapter/out/db";
import { UserMapper } from "../../adapter/in/mapper";

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

  public async createUser(data: CreateUserInput): Promise<boolean> {
    this._logger.info(data, "createUser()");
    try {
      const user = new User({
        firstName: data.firstName,
        lastName: data.lastName,
        emailAddress: new EmailAddress(data?.emailAddress),
        password: data.password,
      });

      await this._userRepository.save(UserMapper.toDTOFromEntity(user));
      return true;
    } catch (err) {
      this._logger.error(err, "createUser()");
      return false;
    }
  }

  public async removeUserById(id: number): Promise<boolean> {
    this._logger.info({ id }, "removeUserById()");
    const userExists = await this.findUserById(id);
    if (!userExists) {
      throw new Error(`Guest with email ${id} doesn't exist`);
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
      return user;
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
      return user;
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
