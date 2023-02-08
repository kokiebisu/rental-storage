import { BaseRestClient } from "./base";

export default class ListingRestClient extends BaseRestClient {
  public async addListing(
    lenderId: string,
    streetAddress: string,
    latitude: number,
    longitude: number,
    imageUrls: string[],
    title: string,
    feeAmount: number,
    feeCurrency: string,
    feeType: string
  ): Promise<{ uid: string } | undefined> {
    try {
      const response = await this.client.post(`/listings`, {
        lenderId,
        streetAddress,
        latitude,
        longitude,
        imageUrls,
        title,
        feeAmount,
        feeCurrency,
        feeType,
      });
      return response.data as { uid: string };
    } catch (err) {
      console.error(err);
    }
  }

  public async findListingById(
    listingId: string
  ): Promise<{ listing: Listing } | undefined> {
    try {
      const response = await this.client.get(`/listings/${listingId}`);
      return response.data;
    } catch (err) {
      console.error(err);
    }
  }

  public async findListingsWithinLatLng(
    latitude: number,
    longitude: number,
    range: number
  ): Promise<{ listings: Listing[] } | undefined> {
    try {
      const response = await this.client.get(
        `/listings?lat=${latitude}&lng=${longitude}&range=${range}`
      );
      return response.data;
    } catch (err) {
      console.error(err);
    }
  }

  public async removeListingById(uid: string): Promise<void> {
    try {
      await this.client.delete(`/listings/${uid}`);
    } catch (err) {
      console.error(err);
    }
  }

  public async findListingsByUserId(
    userId: string
  ): Promise<{ listings: Listing[] } | undefined> {
    try {
      const response = await this.client.get(`/listings?userId=${userId}`);
      return response.data;
    } catch (err) {
      console.error(err);
    }
  }
}
