import { UserInterface } from "../../Types";

export interface UserMessageSender {
  userCreated(data: UserInterface): Promise<void>;
}
