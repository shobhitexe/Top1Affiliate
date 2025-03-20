import { CommissionTxn } from "@/types";
import Image from "next/image";

export default function Commissions({ data }: { data: CommissionTxn[] }) {
  return (
    <div className="bg-white shadow-sm rounded-2xl p-4">
      <div className="font-semibold">Your Commissions</div>
      <div className="flex flex-col gap-4 mt-7">
        {data.map((item) => (
          <CommissionCard
            time={item.date}
            status={"sent"}
            key={item.name}
            {...item}
          />
        ))}
      </div>
    </div>
  );
}

function CommissionCard({
  name,
  time,
  status,
  amount,
}: {
  name: string;
  time: string;
  status: string;
  amount: number;
}) {
  return (
    <div className="flex items-center justify-between">
      <div className="flex items-center gap-4">
        <Image
          src={
            status === "sent"
              ? "/images/dashboard/sent.svg"
              : "/images/dashboard/pending.svg"
          }
          alt={name}
          width={37}
          height={37}
        />
        <div className="flex flex-col text-sm font-semibold">
          <div>{name}</div>
          <div className="text-gray">{time}</div>
        </div>
      </div>

      <div
        className={`${status === "sent" ? "text-[#48BB78]" : ""} font-medium`}
      >
        {status === "sent" ? `$${amount.toLocaleString()}` : "Pending"}
      </div>
    </div>
  );
}
