import { BaseRestClient } from "./base";

export default class SpaceRestClient extends BaseRestClient {
  public async createSpace(
    lenderId: string,
    streetAddress: string,
    latitude: number,
    longitude: number,
    imageUrls: string[],
    title: string,
    feeType: string,
    feeAmount: number,
    feeCurrency: string
  ) {
    return (
      await this.client.post(`/spaces`, {
        lenderId,
        streetAddress,
        latitude,
        longitude,
        imageUrls,
        title,
        feeType,
        feeAmount,
        feeCurrency,
      })
    ).data;
  }

  public async findSpace(id: string) {
    return (await this.client.get(`/spaces/${id}`)).data;
  }

  public async deleteSpace(id: string) {
    return (await this.client.delete(`/spaces/${id}`)).data;
  }

  public async findSpaces(userId: string) {
    return (await this.client.get(`/spaces?userId=${userId}`)).data;
  }
}
