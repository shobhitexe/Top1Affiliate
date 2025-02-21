"use client";

import { Area, AreaChart, CartesianGrid, XAxis, YAxis, Legend } from "recharts";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";

const chartData = [
  { month: "January", commission: 245, withdrawal: 92 },
  { month: "February", commission: 312, withdrawal: 105 },
  { month: "March", commission: 198, withdrawal: 78 },
  { month: "April", commission: 276, withdrawal: 120 },
  { month: "May", commission: 230, withdrawal: 95 },
  { month: "June", commission: 180, withdrawal: 210 },
  { month: "July", commission: 310, withdrawal: 112 },
  { month: "August", commission: 256, withdrawal: 88 },
  { month: "September", commission: 274, withdrawal: 102 },
  { month: "October", commission: 140, withdrawal: 170 },
  { month: "November", commission: 287, withdrawal: 110 },
  { month: "December", commission: 299, withdrawal: 97 },
];

const chartConfig = {
  commission: {
    label: "Commission",
    color: "#4FD1C5",
  },
  withdrawal: {
    label: "Withdrawal",
    color: "#2D3748",
  },
} satisfies ChartConfig;

export default function SalesChart() {
  return (
    <Card>
      <CardHeader>
        <CardTitle>Sales Overview</CardTitle>
      </CardHeader>
      <CardContent className="p-0 sm:-left-3 -left-5 relative">
        <ChartContainer config={chartConfig}>
          <AreaChart
            data={chartData}
            margin={{
              left: 12,
              right: 12,
            }}
          >
            <CartesianGrid vertical={false} />
            <XAxis
              dataKey="month"
              tickLine={false}
              axisLine={false}
              tickMargin={8}
              tickFormatter={(value) => value.slice(0, 3)}
            />
            <YAxis
              tickLine={false}
              axisLine={false}
              tickMargin={8}
              tickFormatter={(value) => `$${value}`}
            />
            <Legend
              verticalAlign="top"
              align="right"
              wrapperStyle={{ marginBottom: 10 }}
            />

            <ChartTooltip cursor={false} content={<ChartTooltipContent />} />
            <defs>
              <linearGradient id="fillCommission" x1="0" y1="0" x2="0" y2="1">
                <stop
                  offset="5%"
                  stopColor="var(--color-commission)"
                  stopOpacity={0.8}
                />
                <stop
                  offset="95%"
                  stopColor="var(--color-commission)"
                  stopOpacity={0.1}
                />
              </linearGradient>
              <linearGradient id="fillWithdrawal" x1="0" y1="0" x2="0" y2="1">
                <stop
                  offset="5%"
                  stopColor="var(--color-withdrawal)"
                  stopOpacity={0.8}
                />
                <stop
                  offset="95%"
                  stopColor="var(--color-withdrawal)"
                  stopOpacity={0.1}
                />
              </linearGradient>
            </defs>
            <Area
              dataKey="commission"
              type="natural"
              fill="url(#fillCommission)"
              stroke="var(--color-commission)"
              stackId="a"
              name="Commission"
            />
            <Area
              dataKey="withdrawal"
              type="natural"
              fill="url(#fillWithdrawal)"
              stroke="var(--color-withdrawal)"
              stackId="a"
              name="Withdrawal"
            />
          </AreaChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}
