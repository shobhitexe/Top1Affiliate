import { DataTable } from "@/components/ui/data-table";
import { BackendURL } from "@/config/env";
import Link from "next/link";
import { buttonVariants } from "@/components/ui/button";
import { listColumns } from "../../list-columns";
import { AffiliatePath } from "@/components";
import { AffiliatePathType } from "@/types";

type Affiliates = {
  id: string;
  name: string;
  affiliateId: "string";
  country: string;
  commission: number;
};

async function GetAffiliates(id: string) {
  try {
    const res = await fetch(`${BackendURL}/api/v1/admin/affiliates?id=${id}`);

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

  const affiliates: Affiliates[] = await GetAffiliates(id as string);

  const path: AffiliatePathType[] = await GetAffiliatePath(id);

  const name = path.filter((item) => item.id === id);

  return (
    <div className="flex flex-col sm:gap-4 gap-2 bg-white p-5 shadow-sm rounded-2xl">
      <div className="flex items-center justify-between">
        <div className="font-semibold text-lg">Affiliates</div>
        <Link
          href={`/list/add?id=${id}&name=${name[0].name}`}
          className={`${buttonVariants({})}`}
        >
          Add
        </Link>
      </div>

      <AffiliatePath path={path} />

      <DataTable columns={listColumns} data={affiliates} />
    </div>
  );
}
