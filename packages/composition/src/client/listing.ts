import { BaseRestClient } from "./base";

export default class ListingRestClient extends BaseRestClient {
  public async createListing(
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
      await this.client.post(`/listings`, {
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

  public async findListing(id: string) {
    return (await this.client.get(`/listings/${id}`)).data;
  }

  public async deleteListing(id: string) {
    return (await this.client.delete(`/listings/${id}`)).data;
  }

  public async findListings(userId: string) {
    return (await this.client.get(`/listings?userId=${userId}`)).data;
  }
}
