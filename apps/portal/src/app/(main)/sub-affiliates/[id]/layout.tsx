import { ViewButtons } from "@/components";
import { ReactNode } from "react";

export default async function layout({
  children,
}: {
  params: Promise<{ id: string }>;
  children: ReactNode;
}) {
  return (
    <div className="flex flex-col sm:gap-4 gap-2 bg-white sm:p-5 p-3 shadow-sm rounded-2xl">
      <div className="flex items-center justify-between">
        <div className="font-semibold text-lg">Sub Affiliates</div>

        <ViewButtons />
      </div>

      {children}
    </div>
  );
}
