import { UserInterface } from "../../types";

export interface UserMessageSender {
  userCreated(data: UserInterface): Promise<void>;
}
