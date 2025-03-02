import { DataTable, weeklyCommissionColumn } from "@/components";
import { DatePickerWithRange } from "@/components/ui/date-picker-range";
import { Input } from "@/components/ui/input";
import { SearchIcon } from "lucide-react";
import { Fragment } from "react";

const commissions = [
  {
    crm_id: "#687451",
    full_name: "Esthera Jackson",
    transaction_date: "20/04/2024",
    country: "United States",
    commission_amount: "$48,684",
    transaction_type: "DEPOSIT",
  },
  {
    crm_id: "#681245",
    full_name: "Alexa Liras",
    transaction_date: "22/04/2024",
    country: "Canada",
    commission_amount: "$32,826",
    transaction_type: "DEPOSIT",
  },
  {
    crm_id: "#687235",
    full_name: "Laurent Michael",
    transaction_date: "26/04/2024",
    country: "France",
    commission_amount: "$12,123",
    transaction_type: "WITHDRAWAL",
  },
  {
    crm_id: "#687852",
    full_name: "Freduardo Hill",
    transaction_date: "30/04/2024",
    country: "United States",
    commission_amount: "$8,684",
    transaction_type: "DEPOSIT",
  },
  {
    crm_id: "#687479",
    full_name: "Daniel Thomas",
    transaction_date: "08/05/2024",
    country: "Egypt",
    commission_amount: "$6,422",
    transaction_type: "WITHDRAWAL",
  },
  {
    crm_id: "#687248",
    full_name: "Mark Michael",
    transaction_date: "09/05/2024",
    country: "India",
    commission_amount: "$5,421",
    transaction_type: "DEPOSIT",
  },
  {
    crm_id: "#687233",
    full_name: "Esthera Jackson",
    transaction_date: "20/05/2024",
    country: "France",
    commission_amount: "$8,653",
    transaction_type: "DEPOSIT",
  },
  {
    crm_id: "#687467",
    full_name: "Alexa",
    transaction_date: "24/05/2024",
    country: "United States",
    commission_amount: "$5,421",
    transaction_type: "WITHDRAWAL",
  },
  {
    crm_id: "#687843",
    full_name: "Laurent Thomas",
    transaction_date: "18/06/2024",
    country: "Executive",
    commission_amount: "$2,652",
    transaction_type: "DEPOSIT",
  },
];

const stats = [
  {
    title: "Registrations for the Week",
    value: "359",
    month: "952 Registrations this month",
  },
  {
    title: "Total Deposits for the Week",
    value: "$15,658.56",
    month: "$28,658 Deposits this month",
  },
  {
    title: "Total Withdrawals for the Week",
    value: "$3,358.24",
    month: "$14,985.65 Withdrawals this month",
  },
  {
    title: "Total Commissions for the Week",
    value: "$13,658.48",
    month: "$44,896.52 Commissions this month",
  },
];

export default function Page() {
  return (
    <div className="flex flex-col gap-4">
      <div className="sm:bg-[#015559] sm:p-5 p-1 rounded-lg md:flex grid sm:grid-cols-2 grid-cols-1 sm:gap-4 gap-2 justify-around items-center">
        {stats.map((item, index) => (
          <Fragment key={item.title}>
            <div className="flex flex-col gap-1 sm:px-5 max-sm:bg-[#015559] max-sm:p-3 max-sm:rounded-md">
              <div className="text-white font-semibold md:text-base text-sm">
                {item.title}
              </div>
              <div className="text-white md:text-4xl text-3xl font-bold">
                {item.value}
              </div>
              <div className="text-[#98EFD4] md:text-sm text-xs">
                {item.month}
              </div>
            </div>

            {index !== stats.length - 1 && (
              <div className="w-[1px] h-16 bg-[#98EFD4] md:flex hidden"></div>
            )}
          </Fragment>
        ))}
      </div>

      <div className="w-full bg-[#E8E8E8] h-px" />

      <div className="flex sm:flex-row flex-col gap-2 sm:items-center justify-between">
        <DatePickerWithRange className="sm:w-fit w-full" />

        <div className="relative">
          <SearchIcon className="absolute top-1/2 -translate-y-1/2 left-2 w-5 h-5" />
          <Input
            className="sm:w-fit w-full bg-white pl-8 h-9"
            placeholder="Search Transaction"
          />
        </div>
      </div>

      <div className="flex flex-col sm:gap-4 gap-2 bg-white p-5 shadow-sm rounded-2xl">
        <div className="font-semibold text-lg">Total Commissions</div>

        <div className="w-full max-w-full overflow-x-auto">
          <DataTable columns={weeklyCommissionColumn} data={commissions} />
        </div>
      </div>
    </div>
  );
}
