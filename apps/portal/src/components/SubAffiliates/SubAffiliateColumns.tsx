"use client";

import { ColumnDef } from "@tanstack/react-table";
import Link from "next/link";
import DetailsDropdown from "./DetailsDropdown";

export const subaffiliateColumns: ColumnDef<unknown>[] = [
  {
    accessorKey: "affiliateId",
    header: "CRM ID",
    cell: ({ row }) => {
      const affiliateId = row.getValue("affiliateId") as string;

      return <DetailsDropdown id={affiliateId} name={`#${affiliateId}`} />;
    },
  },
  {
    accessorKey: "name",
    header: "FIRST NAME",
    cell: ({ row }) => {
      const name = row.getValue("name") as string;
      const id = row.getValue("id") as string;

      return (
        <Link href={`/sub-affiliates/${id}/table`} className={`underline`}>
          {name}
        </Link>
      );
    },
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
