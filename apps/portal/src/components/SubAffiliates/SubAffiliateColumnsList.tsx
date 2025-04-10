"use client";

import { ColumnDef } from "@tanstack/react-table";
import DetailsDropdown from "./DetailsDropdown";

export const subaffiliateColumnsList: ColumnDef<unknown>[] = [
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
    accessorKey: "depth",
    header: "Level",
  },
];
