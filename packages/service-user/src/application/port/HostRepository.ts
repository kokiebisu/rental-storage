import { HostInterface } from "@rental-storage-project/common";

export interface HostRepository {
  setup(): Promise<void>;
  save(data: HostInterface): Promise<{ insertId: string }>;
  delete(hostId: string): Promise<HostInterface>;
  findOneById(id: string): Promise<HostInterface>;
}
