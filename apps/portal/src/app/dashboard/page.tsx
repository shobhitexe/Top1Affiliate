import {
  Commissions,
  MonthlyBarChart,
  ReferralLinks,
  SalesChart,
  TotalStats,
  WeeklyStats,
} from "@/components";

export default function Page() {
  return (
    <div className="flex flex-col gap-4 overflow-hidden sm:px-4 max-w-full">
      <WeeklyStats />

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
          <Commissions />
        </div>
      </div>
    </div>
  );
}
