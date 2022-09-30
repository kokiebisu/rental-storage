export class TimeUtil {
  public static toDate(date: Date) {
    return date.toISOString().slice(0, 10);
  }

  public static toDateTime(date: Date) {
    return date.toISOString().slice(0, 19).replace("T", " ");
  }
}
