import { BaseRestClient } from "./base";

export default class SpaceRestClient extends BaseRestClient {
  public async createSpace(
    lenderId: string,
    imageUrls: string[],
    title: string,
    description: string,
    location: ILocation
  ) {
    return (
      await this.client.post(`/spaces`, {
        lenderId,
        imageUrls,
        title,
        description,
        location,
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
