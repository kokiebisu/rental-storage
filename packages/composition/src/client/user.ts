import { BaseRestClient } from "./base";

export default class UserRestClient extends BaseRestClient {
  public async findUserById(userId: string) {
    const response = await this.client.get(`/users/${userId}`);
    return response.data;
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
}
