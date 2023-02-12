import { BaseRestClient } from "./base";

export default class UserRestClient extends BaseRestClient {
  public async findUser(id: string) {
    return (await this.client.get(`/users/${id}`)).data;
  }

  public async deleteUser(id: string) {
    return (await this.client.delete(`/users/${id}`)).data;
  }

  // WARNING: THIS WILL NOT BE EXPOSED ON APPSYNC
  // PURELY FOR INTEGRATION TEST PURPOSES
  public async createUser(
    emailAddress: string,
    firstName: string,
    lastName: string,
    password: string
  ) {
    return (
      await this.client.post("/users", {
        emailAddress,
        firstName,
        lastName,
        password,
      })
    ).data;
  }
}
