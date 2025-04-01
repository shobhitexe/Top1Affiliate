import { AffiliatePathType } from "@/types";
import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from "@/components/ui/breadcrumb";
import { Home } from "lucide-react";
import { getServerSession } from "next-auth";
import { options } from "@/app/api/auth/[...nextauth]/options";

export default async function AffiliatePath({
  path,
}: {
  path: AffiliatePathType[];
}) {
  const session = await getServerSession(options);

  return (
    <div>
      <Breadcrumb>
        <BreadcrumbList>
          <BreadcrumbItem>
            <BreadcrumbLink href={`/sub-affiliates/${session?.user.id}`}>
              <Home className="w-4 h-4" />
            </BreadcrumbLink>
          </BreadcrumbItem>

          <BreadcrumbSeparator />

          {path.map((item, index) => (
            <BreadcrumbItem key={item.id}>
              {index < path.length - 1 ? (
                <BreadcrumbLink href={`/sub-affiliates/${item.id}`}>
                  {item.name}
                </BreadcrumbLink>
              ) : (
                <BreadcrumbPage>{item.name}</BreadcrumbPage>
              )}
              {index < path.length - 1 && <BreadcrumbSeparator />}
            </BreadcrumbItem>
          ))}
        </BreadcrumbList>
      </Breadcrumb>
    </div>
  );
}
