import { DataTable, leaderboardColumns } from "@/components";

const leaderboard = [
  {
    nickname: "Esthera Jackson",
    country: "United States",
    commissions: "$148,698",
  },
  {
    nickname: "Alexa Liras",
    country: "Canada",
    commissions: "$136,457",
  },
  {
    nickname: "Laurent Michael",
    country: "France",
    commissions: "$98,732",
  },
  {
    nickname: "Freduardo Hill",
    country: "United States",
    commissions: "$84,217",
  },
  {
    nickname: "Daniel Thomas",
    country: "Egypt",
    commissions: "$79,792",
  },
  {
    nickname: "Mark Wilson",
    country: "India",
    commissions: "$73,423",
  },
  {
    nickname: "Esthera Jackson",
    country: "France",
    commissions: "$66,207",
  },
  {
    nickname: "Alexa Liras",
    country: "United States",
    commissions: "$61,982",
  },
  {
    nickname: "Laurent Michael",
    country: "Executive",
    commissions: "$47,326",
  },
  {
    nickname: "Freduardo Hill",
    country: "United States",
    commissions: "$43,172",
  },
  {
    nickname: "Daniel Thomas",
    country: "Czech Republic",
    commissions: "$40,413",
  },
  {
    nickname: "Mark Wilson",
    country: "Canada",
    commissions: "$36,795",
  },
  {
    nickname: "Freduardo Hill",
    country: "Philippines",
    commissions: "$32,264",
  },
];

export default function Page() {
  return (
    <div className="flex flex-col sm:gap-4 gap-2 bg-white p-5 shadow-sm rounded-2xl">
      <div className="font-semibold text-lg">Top 50 Leaderboard</div>

      <div className="w-full max-w-full overflow-x-auto">
        <DataTable columns={leaderboardColumns} data={leaderboard} />
      </div>
    </div>
  );
}
