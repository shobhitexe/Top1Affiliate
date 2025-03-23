"use client";

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
    accessorKey: "id",
    header: "Edit",
    cell: ({ row }) => {
      const id = row.getValue("id") as string;

      return (
        <Link href={`/list/edit/${id}`}>
          <EditIcon />{" "}
        </Link>
      );
    },
  },
];
