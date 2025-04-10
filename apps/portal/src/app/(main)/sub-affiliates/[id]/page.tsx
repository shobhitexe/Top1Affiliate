import {
  AffiliatePath,
  AffiliateTreeView,
  DataTable,
  subaffiliateColumns,
} from "@/components";
import { BackendURL } from "@/config/env";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";

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

  const path = await GetAffiliatePath(id);

  const tree = await GetAffiliateTree(id);

  const affiliates = await GetAffiliates(id);

  return (
    <Tabs defaultValue="tree" className="">
      <TabsList>
        <TabsTrigger value="tree">Tree</TabsTrigger>
        <TabsTrigger value="list">List</TabsTrigger>
        <TabsTrigger value="table">Table</TabsTrigger>
      </TabsList>
      <TabsContent value="tree" className="pt-2">
        <AffiliateTreeView affiliateData={tree} />
      </TabsContent>
      <TabsContent value="list" className="pt-2">
        Make changes to your account here.
      </TabsContent>
      <TabsContent value="table" className="pt-2">
        <AffiliatePath path={path} />
        <DataTable columns={subaffiliateColumns} data={affiliates} />
      </TabsContent>
    </Tabs>
  );
}
