import { Guest, GuestInterface, Host } from "./entity";
import { GuestRepository } from "./repository/guest";
import { isGuest, isHost } from "./utils";

interface UserService {
  registerGuest(data: any): Promise<boolean>;
  registerHost(data: any): Promise<boolean>;
}

export class UserServiceImpl implements UserService {
  private _repository: GuestRepository;
  private constructor(repository: GuestRepository) {
    this._repository = repository;
  }

  public static async create() {
    const repository = await GuestRepository.create();
    return new UserServiceImpl(repository);
  }

  public async registerGuest(data: any): Promise<boolean> {
    if (isGuest(data)) {
      const guest = new Guest(data.firstName, data.lastName);
      // save to guest table
      await this._repository.save(guest.toDTO());
      // send slack message (new user joined!)
      return true;
    } else {
      console.error("Attributes provided didn't match the User interface");
      // send error alert to slack
    }

    return false;
  }

  public async registerHost(data: any): Promise<boolean> {
    if (isHost(data)) {
      const host = new Host(data.firstName, data.lastName);
      // save to guest table
      await this._repository.save(host.toDTO());

      // send slack message (new user joined!)
      return true;
    } else {
      console.error("Attributes provided didn't match the User interface");
      // send error alert to slack
    }

    return false;
  }
}
