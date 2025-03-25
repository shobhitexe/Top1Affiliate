import { EditUser } from "@/components";

import { BackendURL } from "@/config/env";
import { Affiliate } from "@/types";

const affiliate = {
  id: "",
  name: "",
  affiliateId: "",
  country: "",
  commission: 0,
};

async function GetAffiliate(id: string) {
  try {
    const res = await fetch(`${BackendURL}/api/v1/admin/affiliate?id=${id}`);

    if (res.status !== 200) {
      return affiliate;
    }

    const data = await res.json();

    return data.data || affiliate;
  } catch (error) {
    console.log(error);
    return affiliate;
  }
}

export default async function page({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const { id } = await params;

  const affiliate: Affiliate = await GetAffiliate(id as string);

  return <EditUser affiliate={affiliate} id={id} />;
}
