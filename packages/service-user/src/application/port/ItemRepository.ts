import { Item } from "../../domain/Model";

export interface ItemRepository {
  save(data: Item): Promise<Item>;
  delete(id: string): Promise<void>;
  findOneById(id: string): Promise<Item>;
}
