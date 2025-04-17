"use client";

import { ColumnDef } from "@tanstack/react-table";
import DeclinePayout from "./DeclinePayout";
import AcceptPayout from "./AcceptPayout";

export const payoutsColumn: ColumnDef<unknown>[] = [
  {
    accessorKey: "iban",
    header: "#",
    cell: ({ row }) => <div>{row.index + 1}</div>,
  },
  {
    accessorKey: "name",
    header: "Name",
  },
  {
    accessorKey: "affiliateId",
    header: "Affiliate ID",
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
    accessorKey: "walletAddress",
    header: () => <></>,
    cell: () => <></>,
  },
  {
    accessorKey: "amount",
    header: "AMOUNT",
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
    accessorKey: "chainName",
    header: () => <></>,
    cell: () => <></>,
  },
  {
    accessorKey: "type",
    header: "Payout Type",
  },
  {
    accessorKey: "bankName",
    header: () => <></>,
    cell: () => <></>,
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
  {
    accessorKey: "swiftCode",
    header: () => <></>,
    cell: () => <></>,
  },
  {
    accessorKey: "id",
    header: "Action",
    cell: ({ row }) => {
      const status = row.getValue("status") as string;
      const id = row.getValue("id") as string;
      const amount = row.getValue("amount") as number;

      if (status !== "PENDING") {
        return <></>;
      }

      const method = row.getValue("method") as string;

      const iban = row.getValue("iban") as string;
      const swiftCode = row.getValue("swiftCode") as string;
      const bankName = row.getValue("bankName") as string;
      const chainName = row.getValue("chainName") as string;
      const walletAddress = row.getValue("walletAddress") as string;

      return (
        <div className="flex items-center gap-2">
          <AcceptPayout
            id={id}
            amount={amount}
            method={method}
            iban={iban}
            swiftCode={swiftCode}
            bankName={bankName}
            chainName={chainName}
            walletAddress={walletAddress}
          />
          <DeclinePayout id={id} />
        </div>
      );
    },
  },
];
