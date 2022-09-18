import { StreetAddress } from "./street-address";
import { Host } from "./user";

export interface HostListingInfo {
  id: string;
  host: Host;
  emailAddress: string;
  streetAddress: StreetAddress;
}
