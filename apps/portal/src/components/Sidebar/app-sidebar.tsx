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
import PayoutsIcon from "./Icons/payouts";
import CommissionIcon from "./Icons/commission";
import { AnimatePresence, motion } from "framer-motion";
import { useSession } from "next-auth/react";

export function AppSidebar({ ...props }: React.ComponentProps<typeof Sidebar>) {
  const pathname = usePathname();

  const session = useSession();

  const navMain = [
    { title: "Home", url: "/dashboard", icon: HomeIcon },
    { title: "Leaderboard", url: "/leaderboard", icon: LeaderboardIcon },
    { title: "Incentives", url: "#", icon: CardIcon },
    { title: "Statistics", url: "/statistics", icon: StatisticsIcon },
    {
      title: "Weekly Commissions",
      url: "/weekly-commissions",
      icon: CommissionIcon,
    },
    { title: "Payouts", url: "/payouts", icon: PayoutsIcon },
    {
      title: "Sub-Affiliates",
      url: `/sub-affiliates/${session.data?.user.id}/tree`,
      icon: SubIcon,
    },
    { title: "Settings", url: "/settings", icon: SettingsIcon },
  ];

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
              {navMain.map((item) => {
                const isActive = item.url === pathname;

                return (
                  <SidebarMenuItem key={item.title}>
                    <SidebarMenuButton
                      className="h-14 px-4 rounded-2xl"
                      asChild
                    >
                      <Link
                        href={item.url}
                        className="flex items-center gap-2 relative"
                      >
                        <AnimatePresence>
                          {isActive && (
                            <motion.div
                              layoutId="activeMenu"
                              className="absolute inset-0 bg-white border shadow-md rounded-2xl"
                              transition={{ duration: 0.3, ease: "easeInOut" }}
                            />
                          )}
                        </AnimatePresence>

                        <motion.div
                          initial={{ scale: 0.9, opacity: 0.8 }}
                          animate={{ scale: isActive ? 1.1 : 1, opacity: 1 }}
                          exit={{ scale: 0.9, opacity: 0.6 }}
                          transition={{ duration: 0.3 }}
                          className={`relative z-10 ${
                            isActive
                              ? "bg-sidebarBtnBg border"
                              : "bg-white shadow-md border"
                          } rounded-2xl p-2`}
                        >
                          <item.icon
                            fill={isActive ? "white" : "#00987C"}
                            className="w-5 h-5"
                          />
                        </motion.div>

                        <motion.span
                          initial={{ opacity: 0 }}
                          animate={{ opacity: 1 }}
                          exit={{ opacity: 0.5 }}
                          whileTap={{ scale: 0.95 }}
                          transition={{ duration: 0.3 }}
                          className="relative z-10 text-gray font-extrabold font-redhat"
                        >
                          {item.title}
                        </motion.span>
                      </Link>
                    </SidebarMenuButton>
                  </SidebarMenuItem>
                );
              })}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>

        <SidebarGroup>
          <SidebarGroupContent>
            <div className="bg-black text-white p-4 rounded-2xl flex flex-col gap-5 bg-[url('/images/sidebar/Background.png')] bg-cover bg-center">
              <div>
                <div className="font-semibold">Need help?</div>
                {/* <div>Please check our docs</div> */}
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
