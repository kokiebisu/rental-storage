import { Item } from "../../domain/model";

export interface ItemRepository {
  save(data: Item): Promise<Item>;
  delete(id: string): Promise<Item>;
  findOneById(id: string): Promise<Item>;
}
