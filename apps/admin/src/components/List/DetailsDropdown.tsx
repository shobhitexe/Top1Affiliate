"use client";

import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { BackendURL } from "@/config/env";
import fetcher from "@/lib/fetcher";
import { useState } from "react";

import useSWR from "swr";

export default function DetailsDropdown({ id }: { id: string }) {
  const [open, setOpen] = useState(false);

  const { data } = useSWR<{
    data: {
      registrations: number;
      deposits: number;
      withdrawals: number;
      commission: number;
    };
  }>(
    open ? `${BackendURL}/api/v1/data/netstats?affiliateId=${id}` : null,
    fetcher
  );

  return (
    <DropdownMenu open={open} onOpenChange={() => setOpen((prev) => !prev)}>
      <DropdownMenuTrigger className="underline">{id}</DropdownMenuTrigger>
      <DropdownMenuContent>
        <DropdownMenuLabel>Stats</DropdownMenuLabel>
        <DropdownMenuSeparator />
        <DropdownMenuItem>
          Total Registrations: {data?.data.registrations || 0}
        </DropdownMenuItem>
        <DropdownMenuItem>
          Total Deposits: {data?.data.deposits || 0}
        </DropdownMenuItem>
        <DropdownMenuItem>
          Total Withdrawals: {data?.data.withdrawals || 0}
        </DropdownMenuItem>
        <DropdownMenuItem>
          Total Commissions: {data?.data.commission || 0}
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
}
