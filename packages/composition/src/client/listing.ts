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
  ) {
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
      return response.data;
    } catch (err) {
      console.error(err);
    }
  }

  public async findListingById(listingId: string) {
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
  ) {
    try {
      const response = await this.client.get(
        `/listings?lat=${latitude}&lng=${longitude}&range=${range}`
      );
      return response.data;
    } catch (err) {
      console.error(err);
    }
  }

  public async removeListingById(uid: string) {
    try {
      const response = await this.client.delete(`/listings/${uid}`);
      return response;
    } catch (err) {
      console.error(err);
    }
  }

  public async findListingsByUserId(userId: string) {
    try {
      const response = await this.client.get(`/listings?userId=${userId}`);
      return response.data;
    } catch (err) {
      console.error(err);
    }
  }
}
