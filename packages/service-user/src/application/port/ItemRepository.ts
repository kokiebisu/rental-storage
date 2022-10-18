import { Item } from "../../Domain/Model";

export interface ItemRepository {
  save(data: Item): Promise<Item>;
  delete(id: string): Promise<Item>;
  findOneById(id: string): Promise<Item>;
}
