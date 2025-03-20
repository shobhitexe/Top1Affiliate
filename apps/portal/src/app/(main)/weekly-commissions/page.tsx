import { options } from "@/app/api/auth/[...nextauth]/options";
import { DataTable, weeklyCommissionColumn } from "@/components";
import { Input } from "@/components/ui/input";
import { BackendURL } from "@/config/env";
import { SearchIcon } from "lucide-react";
import { getServerSession } from "next-auth";
import { Fragment } from "react";
import DateFilter from "./DateFilter";
import { CommissionTxn, StatsData } from "@/types";

async function GetWeeklyStats(id: string) {
  try {
    const res = await fetch(
      `${BackendURL}/api/v1/data/weekly-stats?affiliateId=${id}`
    );

    if (res.status !== 200) {
      return {
        registrations: 0,
        deposits: 0,
        withdrawals: 0,
        commission: 0,
        registrationsMonthly: 0,
        depositsMonthly: 0,
        withdrawalsMonthly: 0,
        commissionMonthly: 0,
      };
    }

    const data = await res.json();

    return (
      data.data || {
        registrations: 0,
        deposits: 0,
        withdrawals: 0,
        commission: 0,
        registrationsMonthly: 0,
        depositsMonthly: 0,
        withdrawalsMonthly: 0,
        commissionMonthly: 0,
      }
    );
  } catch (error) {
    console.log(error);
    return {
      registrations: 0,
      deposits: 0,
      withdrawals: 0,
      commission: 0,
      registrationsMonthly: 0,
      depositsMonthly: 0,
      withdrawalsMonthly: 0,
      commissionMonthly: 0,
    };
  }
}

async function GetTransactions(id: string, from: string, to: string) {
  try {
    const res = await fetch(
      `${BackendURL}/api/v1/data/transactions?affiliateId=${id}&from=${from}&to=${to}`
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

  const statsdata: StatsData = await GetWeeklyStats(
    session?.user.affiliateId || ""
  );

  const transactions: CommissionTxn[] = await GetTransactions(
    session?.user.affiliateId || "",
    from as string,
    to as string
  );

  const stats = [
    {
      title: "Registrations for the Week",
      value: `${statsdata.registrations}`,
      month: `${statsdata.registrationsMonthly} Registrations this month`,
    },
    {
      title: "Total Deposits for the Week",
      value: `$${statsdata.deposits}`,
      month: `$${statsdata.depositsMonthly} Deposits this month`,
    },
    {
      title: "Total Withdrawals for the Week",
      value: `$${statsdata.withdrawals}`,
      month: `$${statsdata.withdrawalsMonthly} Withdrawals this month`,
    },
    {
      title: "Total Commissions for the Week",
      value: `$${statsdata.commission}`,
      month: `$${statsdata.commissionMonthly} Commissions this month`,
    },
  ];

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
        <div className="font-semibold text-lg">Total Commissions</div>

        <div className="w-full max-w-full overflow-x-auto">
          <DataTable
            columns={weeklyCommissionColumn}
            data={transactions || []}
          />
        </div>
      </div>
    </div>
  );
}
