import { GuestInterface, HostInterface } from "@rental-storage-project/common";
import { Guest, Host } from "./entity";
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

  private constructor(
    guestRepository: GuestRepository,
    hostRepository: HostRepository
  ) {
    this._guestRepository = guestRepository;
    this._hostRepository = hostRepository;
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
      const guest = new Guest(data.firstName, data.lastName);
      // save to guest table
      await this._guestRepository.save(guest.toDTO());
      // send slack message (new user joined!)
      return true;
    } else {
      console.error("Attributes provided didn't match the User interface");
      // send error alert to slack
    }

    return false;
  }

  public async removeGuestById(id: string): Promise<boolean> {
    try {
      await this._guestRepository.delete(id);
      return true;
    } catch (err) {
      console.error(err);
      return false;
    }
  }

  public async findGuestById(id: string): Promise<GuestInterface | null> {
    try {
      const guest = await this._guestRepository.findOneById(id);
      return guest;
    } catch (err) {
      console.error(err);
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
      console.error("Attributes provided didn't match the User interface");
      // send error alert to slack
    }

    return false;
  }

  public async removeHostById(id: string): Promise<boolean> {
    try {
      await this._hostRepository.delete(id);
      return true;
    } catch (err) {
      console.error(err);
      return false;
    }
  }

  public async findHostById(id: string): Promise<HostInterface | null> {
    try {
      const host = await this._hostRepository.findOneById(id);
      return host;
    } catch (err) {
      console.error(err);
      return null;
    }
  }
}
