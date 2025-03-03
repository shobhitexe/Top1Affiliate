"use client";

import { ColumnDef } from "@tanstack/react-table";

export const statisticsColumns: ColumnDef<unknown>[] = [
  {
    accessorKey: "crm_id",
    header: "CRM ID",
    cell: ({ row }) => {
      const id = row.getValue("crm_id") as string;

      return (
        <div className={`${id === "Totals" ? "font-semibold" : ""}`}>{id}</div>
      );
    },
  },
  {
    accessorKey: "first_name",
    header: "FIRST NAME",
  },
  {
    accessorKey: "last_name",
    header: "LAST NAME",
  },
  {
    accessorKey: "registration_date",
    header: "REGISTRATION DATE",
  },
  {
    accessorKey: "country",
    header: "COUNTRY",
  },
  {
    accessorKey: "deposits",
    header: "DEPOSITS",
    cell: ({ row }) => (
      <div
        className={`bg-[#6CA79B] text-white text-center px-4 py-1 rounded-lg`}
      >
        {row.getValue("deposits")}
      </div>
    ),
  },
  {
    accessorKey: "withdrawals",
    header: "WITHDRAWALS",
    cell: ({ row }) => (
      <div
        className={`bg-[#C77D7D] text-white text-center px-4 py-1 rounded-lg`}
      >
        {row.getValue("withdrawals")}
      </div>
    ),
  },
  {
    accessorKey: "net_deposit",
    header: "NET DEPOSIT",
    cell: ({ row }) => (
      <div
        className={`bg-[#019D7F] text-white text-center px-4 py-1 rounded-lg`}
      >
        {row.getValue("net_deposit")}
      </div>
    ),
  },
  {
    accessorKey: "commissions",
    header: "COMMISSIONS",
    cell: ({ row }) => (
      <div
        className={`bg-[#677F89] text-white text-center px-4 py-1 rounded-lg`}
      >
        {row.getValue("commissions")}
      </div>
    ),
  },
];
