import { BookingInterface } from "../../types";

export interface BookingRepository {
  save(booking: BookingInterface): Promise<void>;
  delete(id: string): Promise<any>;
  findById(id: string): Promise<any>;
}
