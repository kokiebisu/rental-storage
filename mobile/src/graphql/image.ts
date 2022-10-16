import gql from "graphql-tag";

export const QUERY_GET_PRESIGNED_URL = gql`
  query MyQuery($filename: String) {
    getPresignedURL(filename: $filename) {
      filename
      url
    }
  }
`;
