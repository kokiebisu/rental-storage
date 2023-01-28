import { BaseRestClient } from "./base";

export default class ListingRestClient extends BaseRestClient {
  public async addListing(
    lenderId: string,
    streetAddress: string,
    latitude: number,
    longitude: number,
    imageUrls: string[],
    title: string,
    fee: unknown
  ) {
    const response = await this.client.post(`/listings`, {
      lenderId,
      streetAddress,
      latitude,
      longitude,
      imageUrls,
      title,
      fee,
    });
    return response.data;
  }

  public async findListingById(listingId: string) {
    const response = await this.client.get(`/listings?uid=${listingId}`);
    return response.data;
  }

  public async findListingsWithinLatLng(
    latitude: number,
    longitude: number,
    range: number
  ) {
    const response = await this.client.get(
      `/listings?lat=${latitude}&lng=${longitude}&range=${range}`
    );
    return response.data;
  }

  public async removeListingById(uid: string) {
    const response = await this.client.delete(`/listings?uid=${uid}`);
    return response.data;
  }

  public async findListingsByUserId(userId: string) {
    const response = await this.client.get(`/listings?userId=${userId}`);
    return response.data;
  }
}
