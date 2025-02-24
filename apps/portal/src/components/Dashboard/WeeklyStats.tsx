import Image from "next/image";

const Stats = [
  {
    title: "Weekly Registrations",
    value: 53000,
    icon: "/images/dashboard/reg.svg",
    change: 55,
  },
  {
    title: "Weekly Deposits",
    value: 2300,
    icon: "/images/dashboard/deposit.svg",
    change: 5,
  },
  {
    title: "Weekly Withdrawals",
    value: 3052,
    icon: "/images/dashboard/withdrawal.svg",
    change: -14,
  },
  {
    title: "Weekly Commissions",
    value: 173000,
    icon: "/images/dashboard/commissions.svg",
    change: 8,
  },
];

export default function WeeklyStats() {
  return (
    <div className="grid lg:grid-cols-4 sm:grid-cols-2 grid-cols-1 gap-3">
      {Stats.map((item) => (
        <StatCard key={item.title} {...item} />
      ))}
    </div>
  );
}

function StatCard({
  title,
  value,
  icon,
  change,
}: {
  title: string;
  value: number;
  icon: string;
  change: number;
}) {
  return (
    <div className="flex items-center w-full justify-between bg-white py-3 px-5 rounded-2xl shadow-sm">
      <div className="flex flex-col gap-1">
        <div className="text-gray text-sm font-semibold">{title}</div>
        <div className="flex items-end gap-2">
          <div className="font-extrabold font-redhat">
            {title !== "Weekly Registrations" && "$"}
            {value.toLocaleString()}
          </div>
          <div
            className={`${
              change > 0 ? "text-[#48BB78]" : "text-[#E53E3E]"
            } font-extrabold`}
          >
            {change > 0 && "+"}
            {change}%
          </div>
        </div>
      </div>

      <Image src={icon} alt={title} width={57} height={57} />
    </div>
  );
}
