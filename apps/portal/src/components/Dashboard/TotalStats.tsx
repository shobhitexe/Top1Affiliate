import Image from "next/image";

const stats = [
  {
    title: "Total Registrations",
    value: "32,984",
    icon: "/images/dashboard/total.svg",
  },
  { title: "FTDS", value: "2,42M", icon: "/images/dashboard/wallet.svg" },
  {
    title: "Conversion Rate %",
    value: "48%",
    icon: "/images/dashboard/conversion.svg",
  },
];

export default function TotalStats() {
  return (
    <div className="flex flex-col gap-3">
      {stats.map((item) => (
        <StatCard key={item.title} {...item} />
      ))}
    </div>
  );
}

function StatCard({
  title,
  value,
  icon,
}: {
  title: string;
  value: string;
  icon: string;
}) {
  return (
    <div className="bg-white rounded-2xl shadow-sm flex flex-col gap-1 p-3 w-full">
      <div className="flex items-center gap-2">
        <Image
          src={icon}
          alt={title}
          width={46}
          height={46}
          className="w-[46px] h-[46px]"
        />
        <div className="text-gray text-sm font-semibold">{title}</div>
      </div>
      <div className="font-semibold text-xl px-3">{value}</div>
    </div>
  );
}
