"use client";

import { ColumnDef } from "@tanstack/react-table";

export const weeklyCommissionColumn: ColumnDef<unknown>[] = [
  {
    accessorKey: "id",
    header: "ID",
  },
  {
    accessorKey: "name",
    header: "FULL NAME",
  },
  {
    accessorKey: "date",
    header: "TRANSACTION DATE",
  },
  {
    accessorKey: "country",
    header: "COUNTRY",
  },
  {
    accessorKey: "amount",
    header: "COMMISSION AMOUNT",
    cell: ({ row }) => {
      const commission = row.getValue("amount") as number;

      return (
        <div className="text-white bg-[#7B7B7B] max-w-[80px] flex justify-center rounded-lg w-full py-0.5 px-5">
          ${commission}
        </div>
      );
    },
  },
  {
    accessorKey: "txnType",
    header: "TRANSACTION TYPE",
    cell: ({ row }) => {
      const type = row.getValue("txnType") as string;

      let bg;

      switch (type) {
        case "Deposit":
          bg = "bg-[#51A796]";
          break;
        case "Withdrawal":
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
