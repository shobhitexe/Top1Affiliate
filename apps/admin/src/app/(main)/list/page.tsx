import { DataTable } from "@/components/ui/data-table";
import { BackendURL } from "@/config/env";
import { listColumns } from "./list-columns";
import Link from "next/link";
import { buttonVariants } from "@/components/ui/button";

type Affiliates = {
  id: string;
  name: string;
  affiliateId: "string";
  country: string;
  commission: number;
};

async function GetAffiliates() {
  try {
    const res = await fetch(`${BackendURL}/api/v1/admin/affiliates`);

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

export default async function page() {
  const affiliates: Affiliates[] = await GetAffiliates();

  return (
    <div className="flex flex-col sm:gap-4 gap-2 bg-white p-5 shadow-sm rounded-2xl">
      <div className="flex items-center justify-between">
        <div className="font-semibold text-lg">Affiliates</div>
        <Link href={`/list/add`} className={`${buttonVariants({})}`}>
          Add
        </Link>
      </div>

      <DataTable columns={listColumns} data={affiliates} />
    </div>
  );
}
