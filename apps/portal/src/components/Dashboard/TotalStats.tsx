import { options } from "@/app/api/auth/[...nextauth]/options";
import { WeeklyStatsData } from "@/types";
import { getServerSession } from "next-auth";
import Image from "next/image";

export default async function TotalStats({
  stats,
}: {
  stats: WeeklyStatsData;
}) {
  const session = await getServerSession(options);

  const Stats = [
    {
      title: "Total Registrations",
      value: `${stats.registrations.toLocaleString()}`,
      icon: "/images/dashboard/total.svg",
    },
    {
      title: "FTDS",
      value: `${stats.ftds.toLocaleString()}`,
      icon: "/images/dashboard/wallet.svg",
    },
    {
      title: "Commission %",
      value: `${session?.user.commission}`,
      icon: "/images/dashboard/conversion.svg",
    },
  ];

  return (
    <div className="flex flex-col gap-3">
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
}: {
  title: string;
  value: string;
  icon: string;
}) {
  return (
    <div className="bg-white rounded-2xl shadow-sm flex flex-col gap-1 p-3 w-full">
      <div className="flex items-center sm:gap-2 gap-1">
        <Image
          src={icon}
          alt={title}
          width={46}
          height={46}
          className="w-[46px] h-[46px]"
        />
        <div className="text-gray text-sm font-semibold">{title}</div>
      </div>
      <div className="text-lg px-3 font-redhat font-extrabold">{value}</div>
    </div>
  );
}
