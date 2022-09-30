import { GuestMapper, HostMapper } from "../../adapter/in/mapper";
import { Guest, Host } from "../../domain/model";
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
        firstName: data.firstName,
        lastName: data.lastName,
      });

      await this._guestRepository.save(GuestMapper.toDTOFromEntity(guest));
      // send slack message (new user joined!)
      return true;
    } catch (err) {
      this._logger.error(err, "registerGuest()");
      return false;
    }
  }

  public async removeGuestById(id: string): Promise<boolean> {
    try {
      await this._guestRepository.delete(id);
      return true;
    } catch (err) {
      this._logger.error(err, "removeGuestById");
      return false;
    }
  }

  public async findGuestById(id: string): Promise<GuestInterface | null> {
    try {
      const guest = await this._guestRepository.findOneById(id);
      return guest;
    } catch (err) {
      this._logger.error(err, "findGuestById()");
      return null;
    }
  }

  public async registerHost(data: any): Promise<boolean> {
    try {
      const host = new Host({
        firstName: data.firstName,
        lastName: data.lastName,
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

  public async removeHostById(id: string): Promise<boolean> {
    try {
      await this._hostRepository.delete(id);
      return true;
    } catch (err) {
      this._logger.error(err, "removeHostById()");
      return false;
    }
  }

  public async findHostById(id: string): Promise<HostInterface | null> {
    try {
      const host = await this._hostRepository.findOneById(id);
      return host;
    } catch (err) {
      this._logger.error(err, "findHostById()");
      return null;
    }
  }

  public async updateStoredItems(booking: BookingInterface): Promise<boolean> {
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
