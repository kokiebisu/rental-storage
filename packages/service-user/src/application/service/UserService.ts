import { EmailAddress, User } from "../../domain/model";
import {
  CreateUserInput,
  PaymentService,
  UserRepository,
  UserService,
} from "../port";

import { LoggerUtil } from "../../utils";
import { UserInterface } from "../../types";
import { UserRepositoryImpl } from "../../adapter/out/db";
import { UserMapper } from "../../adapter/in/mapper";
import { PaymentServiceImpl } from "./PaymentService";

export class UserServiceImpl implements UserService {
  private _userRepository: UserRepository;
  private _paymentService: PaymentService;

  private _logger: LoggerUtil;

  private constructor(
    userRepository: UserRepository,
    paymentService: PaymentService
  ) {
    this._userRepository = userRepository;
    this._paymentService = paymentService;
    this._logger = new LoggerUtil("UserServiceImpl");
  }

  public static async create() {
    const userRepository = await UserRepositoryImpl.create();
    const paymentService = await PaymentServiceImpl.create();
    await userRepository.setup();
    return new UserServiceImpl(userRepository, paymentService);
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

      await this._paymentService.addPayment({
        userId: savedUser.id,
        emailAddress: savedUser.emailAddress.value,
        firstName: savedUser.firstName,
        lastName: savedUser.lastName,
      });

      return UserMapper.toDTOFromEntity(savedUser);
    } catch (err) {
      this._logger.error(err, "createUser()");
      throw err;
    }
  }

  public async removeById(id: number): Promise<boolean> {
    this._logger.info({ id }, "removeUserById()");
    const userExists = await this.findById(id);
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

  public async findById(id: number): Promise<UserInterface | null> {
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

  public async findByEmail(
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

  // public async addItem(data: ItemInterface): Promise<boolean> {
  //   this._logger.info(data, "addItem()");
  //   try {
  //     const entity = new Item({
  //       name: data.name,
  //       imageUrls: data.imageUrls,
  //       ownerId: data.ownerId,
  //       listingId: data.listingId,
  //       createdAt: new Date(data.createdAt),
  //       ...(data.updatedAt && { updatedAt: new Date(data.updatedAt) }),
  //     });
  //     const dto = ItemMapper.toDTOFromEntity(entity);
  //     const result = await this._ItemRepository.save(dto);
  //     console.log("INSERTED RESULT: ", result);
  //     if (result?.insertId) {
  //       await this._broker.dispatchItemSaved({
  //         ...dto,
  //         id: result.insertId,
  //       });
  //     }

  //     return true;
  //   } catch (err) {
  //     this._logger.error("Something went wrong", "addItem()");
  //     return false;
  //   }
  // }
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
