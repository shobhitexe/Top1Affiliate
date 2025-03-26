import { NetStats } from "@/components";
import { BackendURL } from "@/config/env";
import { Fragment } from "react";

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

const stats = {
  registrations: 0,
  deposits: 0,
  withdrawals: 0,
  commission: 0,
};

async function GetNetStats(id: string) {
  try {
    const res = await fetch(
      `${BackendURL}/api/v1/data/netstats?affiliateId=${id}`
    );

    if (res.status !== 200) {
      return stats;
    }

    const data = await res.json();

    return data.data || stats;
  } catch (error) {
    console.log(error);
    return stats;
  }
}

export default async function page({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const { id } = await params;

  const [statsdata, netstats] = await Promise.all([
    GetWeeklyStats(id as string),
    GetNetStats(id as string),
  ]);

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
    <div className="flex flex-col sm:gap-4 gap-5 bg-white sm:p-5 p-2 shadow-sm rounded-2xl">
      <NetStats stats={netstats} />

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
    </div>
  );
}
