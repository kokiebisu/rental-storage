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
    await userRepository.setup();
    const paymentService = await PaymentServiceImpl.create();
    return new UserServiceImpl(userRepository, paymentService);
  }

  public async createUser(data: CreateUserInput): Promise<boolean> {
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

      return true;
    } catch (err) {
      this._logger.error(err, "createUser()");
      return false;
    }
  }

  public async removeById(uid: string): Promise<boolean> {
    this._logger.info({ uid }, "removeUserById()");
    const userExists = await this.findById(uid);
    if (!userExists) {
      throw new Error(`User with email ${uid} doesn't exist`);
    }
    try {
      const user = await this._userRepository.delete(uid);
      return !!user;
    } catch (err) {
      this._logger.error(err, "removeUserById");
      return false;
    }
  }

  public async findById(uid: string): Promise<UserInterface> {
    this._logger.info({ uid }, "findUserById()");
    try {
      const user = await this._userRepository.findOneById(uid);
      if (!user) {
        throw new Error(`User with id ${uid} doesn't exist in db`);
      }
      return UserMapper.toDTOFromEntity(user);
    } catch (err) {
      this._logger.error(err, "findUserById()");
      throw err;
    }
  }

  public async findByEmail(emailAddress: string): Promise<UserInterface> {
    this._logger.info({ emailAddress }, "findUserByEmail()");
    try {
      const user = await this._userRepository.findOneByEmail(emailAddress);
      if (!user) {
        throw new Error(`User with email ${emailAddress} doesn't exist in db`);
      }
      return UserMapper.toDTOFromEntity(user);
    } catch (err) {
      this._logger.error(err, "findUserByEmail()");
      throw err;
    }
  }
}
