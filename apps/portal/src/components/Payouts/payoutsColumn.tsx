"use client";

import { ColumnDef } from "@tanstack/react-table";

export const payoutsColumn: ColumnDef<unknown>[] = [
  {
    accessorKey: "#",
    header: "#",
    cell: ({ row }) => <div>{row.index + 1}</div>,
  },
  {
    accessorKey: "createdAt",
    header: "REQUESTED DATE",
  },
  {
    accessorKey: "method",
    header: "PAYMENT METHOD",
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
    accessorKey: "type",
    header: "Payout Type",
  },
  {
    accessorKey: "status",
    header: "PAYMENT STATUS",
    cell: ({ row }) => {
      const status = row.getValue("status") as string;

      let bg;

      switch (status) {
        case "PAID":
          bg = "bg-[#28806F]";
          break;
        case "PENDING":
          bg = "bg-[#D5A404]";
          break;
        case "REJECTED":
          bg = "bg-[#AC3E3E]";
          break;
      }

      return (
        <div
          className={`text-white ${bg} max-w-[100px] flex justify-center items-center rounded-lg w-full py-0.5 px-5`}
        >
          <span className="relative top-px">{status}</span>
        </div>
      );
    },
  },
];
