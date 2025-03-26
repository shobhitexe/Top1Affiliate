import { options } from "@/app/api/auth/[...nextauth]/options";
import { DataTable, payoutsColumn } from "@/components";
import PayoutsIcon from "@/components/Sidebar/Icons/payouts";
import { Button } from "@/components/ui/button";
import { DatePickerWithRange } from "@/components/ui/date-picker-range";
import { Input } from "@/components/ui/input";
import { BackendURL } from "@/config/env";
import { SearchIcon } from "lucide-react";
import { getServerSession } from "next-auth";
import Image from "next/image";

async function GetBalance(id: string) {
  try {
    const res = await fetch(
      `${BackendURL}/api/v1/data/balance?affiliateId=${id}`
    );

    if (res.status !== 200) {
      return 0;
    }

    const data = await res.json();

    return data.data || 0;
  } catch (error) {
    console.log(error);

    return 0;
  }
}

export default async function Page() {
  const session = await getServerSession(options);

  const balance = await GetBalance(session?.user.affiliateId || "");

  return (
    <div className="flex flex-col gap-4">
      <div className="bg-[#015559] p-5 rounded-lg flex sm:flex-row flex-col sm:gap-0 gap-4 items-center justify-between">
        <div className="text-white flex flex-col sm:gap-3 gap-1 sm:text-left text-center">
          <div className="font-redhat text-lg">Total Available Balance</div>
          <div className="flex items-end gap-1">
            <div className="text-4xl font-bold">${balance}</div>
            {/* <div className="text-[#1EFFCA] sm:flex hidden">+55%</div> */}
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
          <DataTable columns={payoutsColumn} data={[]} />
        </div>
      </div>
    </div>
  );
}
