export interface UserService {
  registerGuest(data: any): Promise<boolean>;
  registerHost(data: any): Promise<boolean>;
}
