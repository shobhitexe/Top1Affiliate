import { AffiliateTreeView } from "@/components";
import { BackendURL } from "@/config/env";

async function GetAffiliateTree(id: string) {
  try {
    const res = await fetch(`${BackendURL}/api/v1/data/tree?id=${id}`);

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

  const tree = await GetAffiliateTree(id);

  return <AffiliateTreeView affiliateData={tree} />;
}
