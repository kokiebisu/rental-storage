import { UserInterface } from "../../domain/types";

export interface UserMessageSender {
  userCreated(data: UserInterface): Promise<void>;
}
