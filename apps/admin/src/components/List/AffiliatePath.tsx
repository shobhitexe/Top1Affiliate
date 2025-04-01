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

export default async function AffiliatePath({
  path,
}: {
  path: AffiliatePathType[];
}) {
  return (
    <div>
      <Breadcrumb>
        <BreadcrumbList>
          <BreadcrumbItem>
            <BreadcrumbLink href={`/list`}>
              <Home className="w-4 h-4" />
            </BreadcrumbLink>
          </BreadcrumbItem>

          <BreadcrumbSeparator />

          {path.map((item, index) => (
            <BreadcrumbItem key={item.id}>
              {index < path.length - 1 ? (
                <BreadcrumbLink href={`/list/sub/${item.id}`}>
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
