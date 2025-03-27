import { ReactNode } from "react";
import PayoutTabs from "./payout-tabs";

export default function layout({ children }: { children: ReactNode }) {
  return (
    <div className="flex flex-col sm:gap-4 gap-5 bg-white sm:p-5 p-4 shadow-sm rounded-2xl">
      <div className="flex items-center justify-between">
        <div className="font-semibold text-lg">Payouts</div>

        <PayoutTabs />
      </div>

      {children}
    </div>
  );
}
