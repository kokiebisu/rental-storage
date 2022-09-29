import { BookingInterface } from "../../types";

export interface BookingRepository {
  save(booking: BookingInterface): Promise<void>;
  delete(id: string): Promise<any>;
  findById(id: string): Promise<any>;
  findAll(): Promise<any>;
  update(id: string, name: string, description: string): Promise<any>;
}
