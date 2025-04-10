"use client";

import { ColumnDef } from "@tanstack/react-table";

export const subaffiliateColumnsList: ColumnDef<unknown>[] = [
  {
    accessorKey: "affiliateId",
    header: "CRM ID",
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
