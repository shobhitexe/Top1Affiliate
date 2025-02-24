"use client";

import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupContent,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarRail,
} from "@/components/ui/sidebar";
import Image from "next/image";
import HomeIcon from "./Icons/home";
import LeaderboardIcon from "./Icons/leaderboard";
import CardIcon from "./Icons/card";
import StatisticsIcon from "./Icons/statistics";
import SubIcon from "./Icons/sub";
import SettingsIcon from "./Icons/settings";
import Link from "next/link";
import { usePathname } from "next/navigation";
import { Button } from "../ui/button";

const navMain = [
  {
    title: "Home",
    url: "/dashboard",
    icon: HomeIcon,
  },
  {
    title: "Leaderboard",
    url: "/leaderboard",
    icon: LeaderboardIcon,
  },
  {
    title: "Incentives",
    url: "#",
    icon: CardIcon,
  },
  {
    title: "Statistics",
    url: "/statistics",
    icon: StatisticsIcon,
  },
  {
    title: "Sub-Affiliates",
    url: "#",
    icon: SubIcon,
  },
  {
    title: "Settings",
    url: "#",
    icon: SettingsIcon,
  },
];

export function AppSidebar({ ...props }: React.ComponentProps<typeof Sidebar>) {
  const pathname = usePathname();

  return (
    <Sidebar {...props}>
      <SidebarHeader>
        <Image
          src={"/images/logo.svg"}
          alt={"logo"}
          width={200}
          height={34}
          className="py-3 mx-auto"
        />

        <div className="bg-gradient-to-r from-transparent via-[#E0E1E2] to-transparent w-full h-[2px]" />
      </SidebarHeader>
      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupContent>
            <SidebarMenu className="gap-2">
              {navMain.map((item) => (
                <SidebarMenuItem key={item.title}>
                  <SidebarMenuButton className="h-14 px-4 rounded-2xl" asChild>
                    <Link
                      href={item.url}
                      className={`flex items-center gap-2 ${
                        item.url === pathname
                          ? "bg-white shadow-md"
                          : "bg-transparent"
                      }`}
                    >
                      <div
                        className={`${
                          item.url === pathname
                            ? "bg-sidebarBtnBg"
                            : "bg-white shadow-sm"
                        } rounded-2xl p-2`}
                      >
                        <item.icon
                          fill={`${
                            item.url === pathname ? "white" : "#00987C"
                          }`}
                          className="w-6 h-6"
                        />
                      </div>

                      <span className={`text-gray font-extrabold font-redhat`}>
                        {item.title}
                      </span>
                    </Link>
                  </SidebarMenuButton>
                </SidebarMenuItem>
              ))}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>

        <SidebarGroup>
          <SidebarGroupContent>
            <div className="bg-black text-white p-4 rounded-2xl flex flex-col gap-5 bg-[url('/images/sidebar/Background.png')] bg-cover bg-center">
              <div>
                <div className="font-semibold">Need help?</div>
                <div>Please check our docs</div>
              </div>

              <Button className="w-full" variant={"secondary"}>
                Support
              </Button>
            </div>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>
      <SidebarRail />
    </Sidebar>
  );
}
