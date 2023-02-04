import { BaseRestClient } from "./base";

export default class UserRestClient extends BaseRestClient {
  public async findUserById(userId: string) {
    const endpoint = `/users/${userId}`;
    try {
      const response = await this.client.get(endpoint);
      return response.data;
    } catch (err) {
      console.error(err);
    }
  }

  public async findByEmail(email: string) {
    const response = await this.client.get(
      `/users/find-by-email?emailAddress=${email}`
    );
    return response.data;
  }

  public async removeUserById(userId: string) {
    const response = await this.client.delete(`/users/${userId}`);
    return response.data;
  }

  // WARNING: THIS WILL NOT BE EXPOSED ON APPSYNC
  // PURELY FOR INTEGRATION TEST PURPOSES
  public async createUser(
    emailAddress: string,
    firstName: string,
    lastName: string,
    password: string
  ) {
    try {
      const response = await this.client.post("/users", {
        emailAddress,
        firstName,
        lastName,
        password,
      });
      return response.data;
    } catch (err) {
      console.error(err);
    }
  }
}
