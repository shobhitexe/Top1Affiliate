"use client";

import { ColumnDef } from "@tanstack/react-table";

export const weeklyCommissionColumn: ColumnDef<unknown>[] = [
  {
    accessorKey: "crm_id",
    header: "CRM ID",
  },
  {
    accessorKey: "full_name",
    header: "FULL NAME",
  },
  {
    accessorKey: "transaction_date",
    header: "TRANSACTION DATE",
  },
  {
    accessorKey: "country",
    header: "COUNTRY",
  },
  {
    accessorKey: "commission_amount",
    header: "COMMISSION AMOUNT",
    cell: ({ row }) => {
      const commission = row.getValue("commission_amount") as number;

      return (
        <div className="text-white bg-[#7B7B7B] max-w-[80px] flex justify-center rounded-lg w-full py-0.5 px-5">
          {commission}
        </div>
      );
    },
  },
  {
    accessorKey: "transaction_type",
    header: "TRANSACTION TYPE",
    cell: ({ row }) => {
      const type = row.getValue("transaction_type") as string;

      let bg;

      switch (type) {
        case "DEPOSIT":
          bg = "bg-[#51A796]";
          break;
        case "WITHDRAWAL":
          bg = "bg-[#C65D5D]";
          break;
      }

      return (
        <div
          className={`text-white ${bg} max-w-[120px] flex justify-center items-center rounded-lg w-full py-0.5 px-5`}
        >
          <span className="relative top-px">{type}</span>
        </div>
      );
    },
  },
];
