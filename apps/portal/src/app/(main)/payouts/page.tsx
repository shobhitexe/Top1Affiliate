import { options } from "@/app/api/auth/[...nextauth]/options";
import {
  DataTable,
  DateFilter,
  payoutsColumn,
  RequestPayoutDialog,
} from "@/components";
import { Input } from "@/components/ui/input";
import { BackendURL } from "@/config/env";
import { SearchIcon } from "lucide-react";
import { getServerSession } from "next-auth";

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

async function GetPayouts(id: string, from: string, to: string) {
  try {
    const res = await fetch(
      `${BackendURL}/api/v1/wallet/payouts?id=${id}&from=${from}&to=${to}`
    );

    if (res.status !== 200) {
      return [];
    }

    const data = await res.json();

    return data.data || [];
  } catch (error) {
    console.log(error);

    return [];
  }
}

export default async function Page({
  searchParams,
}: {
  searchParams: Promise<{ [key: string]: string | string[] | undefined }>;
}) {
  const { from, to } = await searchParams;

  const session = await getServerSession(options);

  const [balance, payouts] = await Promise.all([
    GetBalance(session?.user.affiliateId || ""),
    GetPayouts(session?.user.id || "", from as string, to as string),
  ]);

  return (
    <div className="flex flex-col gap-4">
      <div className="bg-[#015559] p-5 rounded-lg flex sm:flex-row flex-col sm:gap-0 gap-4 items-center justify-between">
        <div className="text-white flex flex-col sm:gap-3 gap-1 sm:text-left text-center">
          <div className="font-redhat text-lg">Total Available Balance</div>
          <div className="sm:flex items-end gap-1">
            <div className="text-4xl font-bold">${balance}</div>
            {/* <div className="text-[#1EFFCA] sm:flex hidden">+55%</div> */}
          </div>
        </div>

        <div className="flex sm:flex-row flex-col items-center gap-3">
          <RequestPayoutDialog balance={balance} type="payout" />
          <RequestPayoutDialog balance={balance} type="transfer" />
        </div>
      </div>

      <div className="w-full bg-[#E8E8E8] h-px" />

      <div className="flex sm:flex-row flex-col gap-2 sm:items-center justify-between">
        <DateFilter />
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
          <DataTable columns={payoutsColumn} data={payouts || []} />
        </div>
      </div>
    </div>
  );
}
