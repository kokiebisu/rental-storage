export default class {
  public static createSpace() {
    return `/spaces`;
  }

  public static findSpace(id: string) {
    return `/spaces/${id}`;
  }

  public static deleteSpace(id: string) {
    return `/spaces/${id}`;
  }

  public static findSpaces({ userId: userId }: { userId: string }) {
    const queryParams = this.buildQueryParams({ userId });
    return `/spaces?${queryParams}`;
  }

  private static buildQueryParams({ userId }: { userId: string }) {
    const queryParams: string[] = [];
    if (userId) {
      queryParams.push(`userId=${userId}`);
    }
    return queryParams.join("&");
  }
}
