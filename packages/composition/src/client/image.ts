export default class ImageResourceURLBuilder {
  public static getPresignedURL(filename: string) {
    return `/images/${filename}`;
  }
}
