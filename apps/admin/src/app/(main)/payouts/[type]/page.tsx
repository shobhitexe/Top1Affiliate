import { payoutsColumn } from "@/components";
import { DataTable } from "@/components/ui/data-table";
import { BackendURL } from "@/config/env";

async function GetPayouts(type: string) {
  try {
    const res = await fetch(
      `${BackendURL}/api/v1/admin/wallet/payouts?type=${type}`
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

export default async function page({
  params,
}: {
  params: Promise<{ type: string }>;
}) {
  const { type } = await params;

  const payouts = await GetPayouts(type);

  return (
    <div>
      <DataTable columns={payoutsColumn} data={payouts || []} />
    </div>
  );
}
