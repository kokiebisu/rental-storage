import { GuestMapper, HostMapper } from "../../adapter/in/mapper";
import { EmailAddress, Guest, Host, Password } from "../../domain/model";
import { GuestRepository, HostRepository, UserService } from "../port";
import { GuestRepositoryImpl, HostRepositoryImpl } from "../../adapter/out/db";
import { LoggerUtil } from "../../utils";
import { BookingInterface, GuestInterface, HostInterface } from "../../types";

export class UserServiceImpl implements UserService {
  private _guestRepository: GuestRepository;
  private _hostRepository: HostRepository;
  private _logger: LoggerUtil;

  private constructor(
    guestRepository: GuestRepository,
    hostRepository: HostRepository
  ) {
    this._guestRepository = guestRepository;
    this._hostRepository = hostRepository;
    this._logger = new LoggerUtil("UserServiceImpl");
  }

  public static async create() {
    const guestRepository = await GuestRepositoryImpl.create();
    const hostRepository = await HostRepositoryImpl.create();

    await guestRepository.setup();
    await hostRepository.setup();
    return new UserServiceImpl(guestRepository, hostRepository);
  }

  public async registerGuest(data: any): Promise<boolean> {
    this._logger.info(data, "registerGuest()");
    try {
      const guest = new Guest({
        firstName: data?.firstName,
        lastName: data?.lastName,
        emailAddress: new EmailAddress(data?.emailAddress),
        password: await Password.create(data.password),
      });

      await this._guestRepository.save(GuestMapper.toDTOFromEntity(guest));
      // send slack message (new user joined!)
      return true;
    } catch (err) {
      this._logger.error(err, "registerGuest()");
      return false;
    }
  }

  public async removeGuestById(id: number): Promise<boolean> {
    this._logger.info({ id }, "removeGuestById()");
    const guestExists = await this.findGuestById(id);
    if (!guestExists) {
      throw new Error(`Guest with email ${id} doesn't exist`);
    }
    try {
      await this._guestRepository.delete(id);
      return true;
    } catch (err) {
      this._logger.error(err, "removeGuestById");
      return false;
    }
  }

  public async findGuestById(id: number): Promise<GuestInterface | null> {
    this._logger.info({ id }, "findGuestById()");
    try {
      const guest = await this._guestRepository.findOneById(id);
      return guest;
    } catch (err) {
      this._logger.error(err, "findGuestById()");
      return null;
    }
  }

  public async findGuestByEmail(
    emailAddress: string
  ): Promise<GuestInterface | null> {
    this._logger.info({ emailAddress }, "findGuestByEmailAddress()");
    try {
      const guest = await this._guestRepository.findOneByEmail(emailAddress);
      return guest;
    } catch (err) {
      this._logger.error(err, "findGuestByEmailAddress()");
      return null;
    }
  }

  public async registerHost(data: any): Promise<boolean> {
    this._logger.info({ data }, "registerHost()");
    try {
      const host = new Host({
        firstName: data.firstName,
        lastName: data.lastName,
        emailAddress: new EmailAddress(data.emailAddress),
        password: await Password.create(data.password),
      });
      // save to guest table
      await this._hostRepository.save(HostMapper.toDTOFromEntity(host));

      // send slack message (new user joined!)
      return true;
    } catch (err: any) {
      // send to DLQ
      this._logger.error(err, "registerHost()");
      return false;
      // send error alert to slack
    }
  }

  public async removeHostById(id: number): Promise<boolean> {
    this._logger.info({ id }, "removeHostById()");
    const hostExists = await this.findHostById(id);
    if (!hostExists) {
      throw new Error(`Host with id ${id} doesn't exist`);
    }
    try {
      await this._hostRepository.delete(id);
      return true;
    } catch (err) {
      this._logger.error(err, "removeHostById()");
      return false;
    }
  }

  public async findHostById(id: number): Promise<HostInterface | null> {
    this._logger.info({ id }, "findHostById()");
    try {
      const host = await this._hostRepository.findOneById(id);
      return host;
    } catch (err) {
      this._logger.error(err, "findHostById()");
      return null;
    }
  }

  public async updateStoredItems(booking: BookingInterface): Promise<boolean> {
    this._logger.info({ booking }, "updatedStoredItems()");
    try {
      await this._guestRepository.updateStoringItem(
        booking.userId,
        booking.items
      );
      return true;
    } catch (err) {
      this._logger.error(err, "updateStoringItem");
      return false;
    }
  }
}
