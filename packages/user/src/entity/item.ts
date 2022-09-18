enum StoringItemSize {
  SMALL = "sm",
  MEDIUM = "md",
  LARGE = "lg",
}

export class StoringItem {
  private _id?: string;
  private name: string;
  private size: StoringItemSize;

  public constructor(name: string, size: StoringItemSize) {
    this.name = name;
    this.size = size;
  }
}
