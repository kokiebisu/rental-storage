import { UserInterface } from "../../domain/types";

export interface UserEventSender {
  userCreated(data: UserInterface): Promise<void>;
}
