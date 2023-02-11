import { BaseRestClient } from "./base";

export default class ListingRestClient extends BaseRestClient {
  public async addListing(
    lenderId: string,
    streetAddress: string,
    latitude: number,
    longitude: number,
    imageUrls: string[],
    title: string,
    feeType: string,
    feeAmount: number,
    feeCurrency: string
  ): Promise<{ uid: string } | undefined> {
    const response = await this.client.post(`/listings`, {
      lenderId,
      streetAddress,
      latitude,
      longitude,
      imageUrls,
      title,
      feeType,
      feeAmount,
      feeCurrency,
    });
    return response.data;
  }

  public async findListingById(
    listingId: string
  ): Promise<{ listing: Listing } | undefined> {
    return (await this.client.get(`/listings/${listingId}`)).data;
  }

  public async removeListingById(listingId: string): Promise<{ uid: string }> {
    return (await this.client.delete(`/listings/${listingId}`)).data;
  }

  public async findListingsWithinLatLng(
    latitude: number,
    longitude: number,
    range: number
  ): Promise<{ listings: Listing[] } | undefined> {
    return (
      await this.client.get(
        `/listings?lat=${latitude}&lng=${longitude}&range=${range}`
      )
    ).data;
  }

  public async findListingsByUserId(
    userId: string
  ): Promise<{ listings: Listing[] }> {
    return (await this.client.get(`/listings?userId=${userId}`)).data;
  }
}
