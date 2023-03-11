import { useState } from "react";
import { Group } from "@mantine/core";
import { DatePicker } from "@mantine/dates";

const Calendar = () => {
  const [value, setValue] = useState<[Date | null, Date | null]>([null, null]);
  return (
    <DatePicker
      type="range"
      numberOfColumns={2}
      value={value}
      onChange={setValue}
    />
  );
};

export default Calendar;
