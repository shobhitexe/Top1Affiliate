import { DataTable, payoutsColumn } from "@/components";
import PayoutsIcon from "@/components/Sidebar/Icons/payouts";
import { Button } from "@/components/ui/button";
import { DatePickerWithRange } from "@/components/ui/date-picker-range";
import { Input } from "@/components/ui/input";
import { SearchIcon } from "lucide-react";
import Image from "next/image";

const payouts = [
  {
    affiliate_name: "Esthera Jackson",
    requested_date: "20/04/2024",
    payment_method: "Crypto",
    commission_amount: "$48,684",
    payment_status: "PAID",
  },
  {
    affiliate_name: "Alexa Liras",
    requested_date: "22/04/2024",
    payment_method: "Wise",
    commission_amount: "$32,826",
    payment_status: "PAID",
  },
  {
    affiliate_name: "Laurent Michael",
    requested_date: "26/04/2024",
    payment_method: "Wire Transfer",
    commission_amount: "$12,123",
    payment_status: "PENDING",
  },
  {
    affiliate_name: "Freduardo Hill",
    requested_date: "30/04/2024",
    payment_method: "Crypto",
    commission_amount: "$8,684",
    payment_status: "PAID",
  },
  {
    affiliate_name: "Daniel Thomas",
    requested_date: "08/05/2024",
    payment_method: "Wire Transfer",
    commission_amount: "$6,422",
    payment_status: "REJECTED",
  },
  {
    affiliate_name: "Mark Michael",
    requested_date: "09/05/2024",
    payment_method: "Crypto",
    commission_amount: "$5,421",
    payment_status: "PAID",
  },
  {
    affiliate_name: "Esthera Jackson",
    requested_date: "20/05/2024",
    payment_method: "Crypto",
    commission_amount: "$8,653",
    payment_status: "PAID",
  },
  {
    affiliate_name: "Alexa",
    requested_date: "24/05/2024",
    payment_method: "Wise",
    commission_amount: "$5,421",
    payment_status: "REJECTED",
  },
  {
    affiliate_name: "Laurent Thomas",
    requested_date: "18/06/2024",
    payment_method: "Crypto",
    commission_amount: "$2,652",
    payment_status: "PENDING",
  },
  {
    affiliate_name: "Laurent Thomas",
    requested_date: "18/06/2024",
    payment_method: "Wise",
    commission_amount: "$2,652",
    payment_status: "PAID",
  },
];

export default function Page() {
  return (
    <div className="flex flex-col gap-4">
      <div className="bg-[#015559] p-5 rounded-lg flex sm:flex-row flex-col sm:gap-0 gap-4 items-center justify-between">
        <div className="text-white flex flex-col sm:gap-3 gap-1 sm:text-left text-center">
          <div className="font-redhat text-lg">Total Available Balance</div>
          <div className="flex items-end gap-1">
            <div className="text-4xl font-bold">$15,658.56</div>
            <div className="text-[#1EFFCA] sm:flex hidden">+55%</div>
          </div>
        </div>

        <div className="flex sm:flex-row flex-col items-center gap-3">
          <Button size={"lg"} className="font-semibold rounded-xl">
            <PayoutsIcon fill="white" />{" "}
            <span className="relative top-px">Request Payout</span>
          </Button>
          <Button size={"lg"} className="font-semibold bg-[#237C81] rounded-xl">
            <Image
              src={"/images/transfer.svg"}
              alt={"transfer"}
              width={22}
              height={20}
            />{" "}
            <span className="relative top-px">Transfer to Trading Acc</span>
          </Button>
        </div>
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
        <div className="font-semibold text-lg">Total Payouts</div>

        <div className="w-full max-w-full overflow-x-auto">
          <DataTable columns={payoutsColumn} data={payouts} />
        </div>
      </div>
    </div>
  );
}
