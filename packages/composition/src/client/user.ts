import { BaseRestClient } from "./base";

export class UserRestClient extends BaseRestClient {
  public async findUserById(userId: string) {
    console.log("FIND USER BY ID: ", userId);
    const response = await this.client.get(`/users/${userId}`);
    console.log("FINDUSERBYID RESPONSE: ", response);
    return response.data;
  }

  public async findByEmail(email: string) {
    const response = await this.client.get(
      `/users/find-by-email/emailAddress=${email}`
    );
    return response.data;
  }

  public async removeUserById(userId: string) {
    const response = await this.client.delete(`/users/${userId}`);
    return response.data;
  }
}
