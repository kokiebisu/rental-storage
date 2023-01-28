import { BaseRestClient } from "./base";

export default class ImageRestClient extends BaseRestClient {
  public async getPresignedURL(filename: string) {
    const response = await this.client.post(`/images`, {
      filename,
    });
    return response.data;
  }
}
