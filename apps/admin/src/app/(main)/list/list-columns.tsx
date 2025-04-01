"use client";

import { DetailsDropdown, SwitchBlockStatus } from "@/components";
import { ColumnDef } from "@tanstack/react-table";
import { EditIcon, User } from "lucide-react";
import Link from "next/link";

export const listColumns: ColumnDef<unknown>[] = [
  {
    accessorKey: "affiliateId",
    header: "Affiliate Id",
    cell: ({ row }) => {
      const affiliateId = row.getValue("affiliateId") as string;

      return (
        <div className="underline cursor-pointer">
          <DetailsDropdown id={affiliateId} />
        </div>
      );
    },
  },
  {
    accessorKey: "name",
    header: "Name",
    cell: ({ row }) => {
      const id = row.getValue("id") as string;
      const name = row.getValue("name") as string;

      return (
        <Link href={`/list/sub/${id}`} className="underline whitespace-nowrap">
          {name}
        </Link>
      );
    },
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
      const affiliateid = row.getValue("affiliateId") as string;

      const blocked = row.getValue("blocked") as boolean;

      return (
        <div className="flex items-center gap-3">
          <Link
            href={`/profile/${affiliateid}`}
            className="group hover:bg-main duration-300 p-1 rounded-full"
          >
            <User className="group-hover:text-white duration-300" />
          </Link>

          <Link
            href={`/list/edit/${id}`}
            className="group hover:bg-main duration-300 p-1.5 rounded-full"
          >
            <EditIcon className="group-hover:text-white duration-300" />
          </Link>

          <SwitchBlockStatus status={blocked} id={id} />
        </div>
      );
    },
  },
];
