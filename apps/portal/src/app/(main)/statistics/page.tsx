import { options } from "@/app/api/auth/[...nextauth]/options";
import { DataTable, statisticsColumns } from "@/components";
import { BackendURL } from "@/config/env";
import { Leads } from "@/types";
import { getServerSession } from "next-auth";

async function GetLeads(affiliateId: string) {
  try {
    const res = await fetch(
      `${BackendURL}/api/v1/data/statistics?affiliateId=${affiliateId}`,
      {
        method: "GET",
      }
    );

    const data = await res.json();

    return data.data || [];
  } catch (error) {
    console.log(error);

    return [];
  }
}

export default async function page() {
  const session = await getServerSession(options);

  const leads: Leads[] = await GetLeads(session?.user.affiliateId || "");

  return (
    <div className="flex flex-col sm:gap-4 gap-2 bg-white p-5 shadow-sm rounded-2xl">
      <div className="font-semibold text-lg">Statistics</div>

      <div className="w-full max-w-full overflow-x-auto">
        <DataTable columns={statisticsColumns} data={leads} />
      </div>
    </div>
  );
}
