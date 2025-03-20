import { DataTable, leaderboardColumns } from "@/components";
import { BackendURL } from "@/config/env";

async function GetDashboardStats() {
  try {
    const res = await fetch(`${BackendURL}/api/v1/data/leaderboard`);

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

export default async function Page() {
  const leaderboard = await GetDashboardStats();

  return (
    <div className="flex flex-col sm:gap-4 gap-2 bg-white p-5 shadow-sm rounded-2xl">
      <div className="font-semibold text-lg">Top 50 Leaderboard</div>

      <div className="w-full max-w-full overflow-x-auto">
        <DataTable columns={leaderboardColumns} data={leaderboard} />
      </div>
    </div>
  );
}
