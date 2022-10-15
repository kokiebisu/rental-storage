import {
  CreateUserInput,
  PaymentService,
  UserEventSender,
  UserMessageSender,
  UserRepository,
  UserService,
} from "../Port";
import { UserInterface } from "../../types";
import { UserRepositoryImpl } from "../../adapter/Repository";
import { EmailAddress, Name, NameType, User } from "../../domain/Model";
import { LoggerUtil } from "../../utils";
import { PaymentServiceImpl } from "./PaymentService";
import { UserMapper } from "../../adapter/Mapper";
import { UserSNSMessageSender } from "../../adapter/MessageSender/UserSNSMessageSender";
import { AWSRegion } from "../../domain/Enum";
import { UserKinesisStreamEventSender } from "../../adapter/EventSender";

export class UserServiceImpl implements UserService {
  private _userRepository: UserRepository;
  private _paymentService: PaymentService;
  private _userMessageSender: UserMessageSender
  private _userEventSender: UserEventSender

  private _logger: LoggerUtil;

  private constructor(
    userRepository: UserRepository,
    paymentService: PaymentService,
    userMessageSender: UserMessageSender,
    userEventSender: UserEventSender
  ) {
    this._userRepository = userRepository;
    this._paymentService = paymentService;
    this._userMessageSender = userMessageSender
    this._userEventSender = userEventSender
    this._logger = new LoggerUtil("UserServiceImpl");
  }

  public static async create() {
    const userRepository = await UserRepositoryImpl.create();
    await userRepository.setup();
    const paymentService = await PaymentServiceImpl.create();
    const userMessageSender = await UserSNSMessageSender.create(AWSRegion.US_EAST_1)
    const userEventSender = await UserKinesisStreamEventSender.create(AWSRegion.US_EAST_1)
    return new UserServiceImpl(userRepository, paymentService, userMessageSender, userEventSender);
  }

  public async createUser(data: CreateUserInput): Promise<UserInterface> {
    this._logger.info(data, "createUser()");
    try {
      let user = new User({
        firstName: new Name(NameType.FirstName, data.firstName),
        lastName: new Name(NameType.LastName, data.lastName),
        emailAddress: new EmailAddress(data?.emailAddress),
        password: data.password,
      });

      const savedUser = await this._userRepository.save(user);

      await this._paymentService.addPayment({
        userId: savedUser.id,
        emailAddress: savedUser.emailAddress.value,
        firstName: savedUser.firstName.value,
        lastName: savedUser.lastName.value,
      });

      const userDTO = UserMapper.toDTOFromEntity(user)
      // await this._userMessageSender.userCreated(userDTO)
      await this._userEventSender.userCreated(userDTO)
      return userDTO
    } catch (err) {
      this._logger.error(err, "createUser()");
      throw err
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
