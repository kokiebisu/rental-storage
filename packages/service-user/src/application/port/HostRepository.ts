import { HostInterface } from "../../types";

export interface HostRepository {
  setup(): Promise<void>;
  save(data: HostInterface): Promise<{ insertId: string }>;
  delete(hostId: string): Promise<HostInterface>;
  findOneById(id: string): Promise<HostInterface>;
}
