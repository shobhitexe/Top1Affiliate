"use client";

import { TrendingUp } from "lucide-react";
import { Bar, BarChart, CartesianGrid, XAxis, YAxis, Legend } from "recharts";

import { Card, CardContent, CardFooter } from "@/components/ui/card";
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";

const chartData = [
  { month: "January", commission: 245 },
  { month: "February", commission: 312 },
  { month: "March", commission: 198 },
  { month: "April", commission: 276 },
  { month: "May", commission: 230 },
  { month: "June", commission: 289 },
  { month: "July", commission: 315 },
  { month: "August", commission: 190 },
  { month: "September", commission: 278 },
  { month: "October", commission: 260 },
  { month: "November", commission: 300 },
  { month: "December", commission: 225 },
];

const chartConfig = {
  commission: {
    label: "Commission",
    color: "hsl(var(--chart-1))",
  },
} satisfies ChartConfig;

export function MonthlyBarChart() {
  return (
    <Card className="shadow-sm border-none">
      <CardContent className="mt-5">
        <ChartContainer
          config={chartConfig}
          className="px-4 rounded-2xl"
          style={{
            background:
              "linear-gradient(81.62deg, #01A180 2.25%, #014250 79.87%)",
          }}
        >
          <BarChart
            className="mt-5 max-sm:-left-5 relative w-full"
            accessibilityLayer
            data={chartData}
            margin={{ left: 12, right: 12 }}
          >
            <CartesianGrid vertical={false} />
            <XAxis
              dataKey="month"
              tickLine={false}
              tickMargin={10}
              axisLine={false}
              tickFormatter={(value) => value.slice(0, 3)}
            />
            <YAxis
              tickLine={false}
              axisLine={false}
              tickMargin={10}
              tickFormatter={(value) => value.toLocaleString()}
              className="text-white"
              style={{ fill: "white" }}
            />
            <Legend
              verticalAlign="top"
              align="right"
              wrapperStyle={{ marginBottom: 10 }}
            />
            <ChartTooltip
              cursor={false}
              content={<ChartTooltipContent hideLabel />}
            />
            <Bar
              dataKey="commission"
              fill="#FFFFFF"
              radius={[5, 5, 0, 0]}
              name="Commission"
            />
          </BarChart>
        </ChartContainer>
      </CardContent>
      <CardFooter className="flex-col items-start gap-2 text-sm">
        <div className="flex gap-2 font-medium leading-none mt-5">
          Monthly Overview <TrendingUp className="h-4 w-4" />
        </div>
        <div
          className="leading-none text-muted-foreground
        "
        >
          (<span className="text-[#48BB78]">+23</span>) than last Month
        </div>
      </CardFooter>
    </Card>
  );
}
