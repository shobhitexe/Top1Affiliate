"use client";

import { DatePickerWithRange } from "@/components/ui/date-picker-range";
import { addDays, endOfDay, startOfDay } from "date-fns";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";
import { DateRange } from "react-day-picker";

export default function DateFilter() {
  const router = useRouter();

  const [date, setDate] = useState<DateRange | undefined>({
    from: addDays(new Date(), -15),
    to: new Date(),
  });

  useEffect(() => {
    const startDate = startOfDay(date?.from || new Date());

    const endDate = endOfDay(date?.to || new Date());

    const queryParams = new URLSearchParams({
      from: startDate.toISOString(),
      to: endDate.toISOString(),
    }).toString();

    router.push(`?${queryParams}`);
  }, [date, router]);

  return (
    <DatePickerWithRange
      date={date}
      setDate={setDate}
      className="sm:w-fit w-full"
    />
  );
}
