import { options } from "@/app/api/auth/[...nextauth]/options";
import {
  Commissions,
  MonthlyBarChart,
  ReferralLinks,
  SalesChart,
  TotalStats,
  WeeklyStats,
} from "@/components";
import { BackendURL } from "@/config/env";
import { DashboardStats } from "@/types";
import { getServerSession } from "next-auth";

const stats = {
  weekly: {
    registrations: 0,
    deposits: 0,
    withdrawals: 0,
    commission: 0,
  },
};

async function GetDashboardStats(id: string) {
  try {
    const res = await fetch(
      `${BackendURL}/api/v1/data/dashboard?affiliateId=${id}`
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

export default async function Page() {
  const session = await getServerSession(options);

  const stats: DashboardStats = await GetDashboardStats(
    session?.user.affiliateId || ""
  );

  return (
    <div className="flex flex-col gap-4 overflow-hidden sm:px-4 max-w-full">
      <WeeklyStats stats={stats.weekly} />

      <div className="grid md:grid-cols-[55%_45%] grid-cols-1 gap-4">
        <div className="flex flex-col gap-4">
          <SalesChart />
          <div className="grid lg:grid-cols-[55%_45%] grid-cols-1 gap-2">
            <div className="bg-white shadow-sm rounded-2xl overflow-hidden">
              <MonthlyBarChart />
            </div>
            <TotalStats />
          </div>
        </div>

        <div className="flex flex-col gap-4">
          <ReferralLinks />
          <Commissions data={stats.commissions || []} />
        </div>
      </div>
    </div>
  );
}
