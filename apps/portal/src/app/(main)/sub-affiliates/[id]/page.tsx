import { AffiliatePath, DataTable, subaffiliateColumns } from "@/components";
import { BackendURL } from "@/config/env";

async function GetAffiliates(id: string) {
  try {
    const res = await fetch(`${BackendURL}/api/v1/data/sub?id=${id}`);

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

async function GetAffiliatePath(id: string) {
  try {
    const res = await fetch(`${BackendURL}/api/v1/data/path?id=${id}`);

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

export default async function page({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const { id } = await params;

  const path = await GetAffiliatePath(id);

  const affiliates = await GetAffiliates(id);

  return (
    <div className="flex flex-col sm:gap-4 gap-2 bg-white p-5 shadow-sm rounded-2xl">
      <div className="flex items-center justify-between">
        <div className="font-semibold text-lg">Sub Affiliates</div>
      </div>

      <AffiliatePath path={path} />

      <DataTable columns={subaffiliateColumns} data={affiliates} />
    </div>
  );
}
