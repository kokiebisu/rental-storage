import { UserInterface } from "../../types";

export interface UserEventSender {
  userCreated(data: UserInterface): Promise<void>;
}
