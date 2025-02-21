import Image from "next/image";

const commissions = [
  {
    name: "Thomas Dany",
    time: "27 March 2020, at 12:30 PM",
    status: "sent",
    amount: 2982,
  },
  {
    name: "Mary S",
    time: "27 March 2020, at 12:30 PM",
    status: "",
    amount: 4658,
  },
  {
    name: "Josue Atkinson",
    time: "26 March 2020, at 13:45 PM",
    status: "sent",
    amount: 1963,
  },
  {
    name: "Jad Thompson",
    time: "26 March 2020, at 12:30 PM",
    status: "sent",
    amount: 3247,
  },
  {
    name: "Cannon Oliver",
    time: "26 March 2020, at 05:00 AM",
    status: "pending",
    amount: 0,
  },
];

export default function Commissions() {
  return (
    <div className="bg-white shadow-sm rounded-2xl p-4">
      <div className="font-semibold">Your Commissions</div>
      <div className="flex flex-col gap-4 mt-7">
        {commissions.map((item) => (
          <CommissionCard key={item.name} {...item} />
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
