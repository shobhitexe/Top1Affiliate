"use client";

import { ColumnDef } from "@tanstack/react-table";
import Link from "next/link";

export const subaffiliateColumns: ColumnDef<unknown>[] = [
  {
    accessorKey: "affiliateId",
    header: "CRM ID",
    cell: ({ row }) => {
      const affiliateId = row.getValue("affiliateId") as string;
      const id = row.getValue("id") as string;

      return (
        <Link href={`/sub-affiliates/${id}/table`} className={`underline`}>
          #{affiliateId}
        </Link>
      );
    },
  },
  {
    accessorKey: "name",
    header: "FIRST NAME",
  },

  {
    accessorKey: "commission",
    header: "Commission %",
  },
  {
    accessorKey: "country",
    header: "COUNTRY",
  },
  {
    accessorKey: "id",
    header: "",
    cell: () => <></>,
  },
];
