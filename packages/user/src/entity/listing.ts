import { Host } from "./host";
import { StreetAddress } from "./street-address";

export interface HostListingInfo {
  id: string;
  host: Host;
  emailAddress: string;
  streetAddress: StreetAddress;
}
