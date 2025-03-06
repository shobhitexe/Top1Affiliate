import { options } from "@/app/api/auth/[...nextauth]/options";
import { DataTable, statisticsColumns } from "@/components";
import { Leads } from "@/types";
import { getServerSession } from "next-auth";

async function GetLeads(cookie: string) {
  try {
    const res = await fetch(
      `https://publicapi.fxlvls.com/management/leads?limit=50&minRegistrationDate=2020-01-01`,
      {
        method: "GET",
        headers: { Cookie: cookie },
      }
    );

    const data = await res.json();

    return data;
  } catch (error) {
    console.log(error);

    return [];
  }
}

export default async function page() {
  const session = await getServerSession(options);

  const leads: Leads[] = await GetLeads(session?.user.cookie || "");

  return (
    <div className="flex flex-col sm:gap-4 gap-2 bg-white p-5 shadow-sm rounded-2xl">
      <div className="font-semibold text-lg">Statistics</div>

      <div className="w-full max-w-full overflow-x-auto">
        <DataTable columns={statisticsColumns} data={leads} />
      </div>
    </div>
  );
}
