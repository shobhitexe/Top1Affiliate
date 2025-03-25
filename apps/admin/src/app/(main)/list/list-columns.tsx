"use client";

import { SwitchBlockStatus } from "@/components";
import { ColumnDef } from "@tanstack/react-table";
import { EditIcon } from "lucide-react";
import Link from "next/link";

export const listColumns: ColumnDef<unknown>[] = [
  {
    accessorKey: "affiliateId",
    header: "Affiliate Id",
  },
  {
    accessorKey: "name",
    header: "Name",
  },
  {
    accessorKey: "country",
    header: "Country",
  },
  {
    accessorKey: "commission",
    header: "Commission %",
  },
  {
    accessorKey: "blocked",
    header: "",
    cell: () => <></>,
  },
  {
    accessorKey: "id",
    header: "Actions",
    cell: ({ row }) => {
      const id = row.getValue("id") as string;

      const blocked = row.getValue("blocked") as boolean;

      return (
        <div className="flex items-center gap-5">
          <Link href={`/list/edit/${id}`}>
            <EditIcon />
          </Link>

          <SwitchBlockStatus status={blocked} id={id} />
        </div>
      );
    },
  },
];
