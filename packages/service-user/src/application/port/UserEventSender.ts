import { UserInterface } from "../../Types";

export interface UserEventSender {
  userCreated(data: UserInterface): Promise<void>;
}
