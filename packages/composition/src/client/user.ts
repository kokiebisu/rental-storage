export default class {
  public static findUser(id: string) {
    return `/users/${id}`;
  }

  public static deleteUser(id: string) {
    return `/users/${id}`;
  }

  // WARNING: THIS WILL NOT BE EXPOSED ON APPSYNC
  // PURELY FOR INTEGRATION TEST PURPOSES
  public static createUser() {
    return `/users`;
  }
}
