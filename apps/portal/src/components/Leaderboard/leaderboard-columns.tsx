"use client";

import { ColumnDef } from "@tanstack/react-table";
import Image from "next/image";

export const leaderboardColumns: ColumnDef<unknown>[] = [
  {
    accessorKey: "ranking",
    header: "RANKING",
    cell: ({ row }) => (
      <div
        className={`${
          row.index === 0
            ? "bg-[#FFB236]"
            : row.index === 1
            ? "bg-[#D7D7D7]"
            : row.index === 2
            ? "bg-[#C77D7D]"
            : "bg-main"
        } w-fit text-white px-4 py-1.5 rounded-xl font-semibold`}
      >
        #{row.index + 1}
      </div>
    ),
  },
  {
    accessorKey: "nickname",
    header: "NICKNAME",
    cell: ({ row }) => {
      const nickname = row.getValue("nickname") as string;

      return (
        <div className="flex items-center w-full gap-2">
          <Image
            src={"/images/leaderboard/Image.svg"}
            alt={"pfp"}
            width={53}
            height={53}
            className="flex-shrink-0"
          />

          <div className="flex-1 min-w-[150px]">{nickname}</div>
        </div>
      );
    },
  },
  {
    accessorKey: "country",
    header: "COUNTRY",
    size: 100,
  },
  {
    accessorKey: "commissions",
    header: "TOTAL COMMISSIONS",
    size: 150,
  },
];
