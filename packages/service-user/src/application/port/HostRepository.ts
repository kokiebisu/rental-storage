import { HostInterface } from "../../types";

export interface HostRepository {
  setup(): Promise<void>;
  save(data: HostInterface): Promise<{ insertId: string }>;
  delete(id: number): Promise<HostInterface>;
  findOneById(id: number): Promise<HostInterface>;
}
