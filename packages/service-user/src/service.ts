import {
  BookingInterface,
  GuestInterface,
  HostInterface,
  LoggerService,
} from "@rental-storage-project/common";
import { Guest, Host } from "./entity";
import { GuestMapper } from "./mapper";
import { HostRepository } from "./repository";
import { GuestRepository } from "./repository/guest";
import { isGuest, isHost } from "./utils";

interface UserService {
  registerGuest(data: any): Promise<boolean>;
  registerHost(data: any): Promise<boolean>;
}

export class UserServiceImpl implements UserService {
  private _guestRepository: GuestRepository;
  private _hostRepository: HostRepository;
  private _logger: LoggerService;

  private constructor(
    guestRepository: GuestRepository,
    hostRepository: HostRepository
  ) {
    this._guestRepository = guestRepository;
    this._hostRepository = hostRepository;
    this._logger = new LoggerService("UserServiceImpl");
  }

  public static async create() {
    const guestRepository = await GuestRepository.create();
    const hostRepository = await HostRepository.create();

    await guestRepository.setup();
    await hostRepository.setup();
    return new UserServiceImpl(guestRepository, hostRepository);
  }

  public async registerGuest(data: any): Promise<boolean> {
    if (isGuest(data)) {
      const guest = new Guest({
        firstName: data.firstName,
        lastName: data.lastName,
      });
      // save to guest table

      await this._guestRepository.save(GuestMapper.toDTOFromEntity(guest));
      // send slack message (new user joined!)
      return true;
    } else {
      this._logger.error(
        "Attributes provided didn't match the User interface",
        "registerGuest"
      );
      // send error alert to slack
    }

    return false;
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
    if (isHost(data)) {
      const host = new Host(data.firstName, data.lastName);
      // save to guest table
      await this._hostRepository.save(host.toDTO());

      // send slack message (new user joined!)
      return true;
    } else {
      this._logger.error(
        "Attributes provided didn't match the User interface",
        "registerHost()"
      );
      // send error alert to slack
    }

    return false;
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
