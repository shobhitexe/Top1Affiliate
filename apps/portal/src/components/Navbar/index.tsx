import { SearchIcon, User2 } from "lucide-react";
import { SidebarTrigger } from "../ui/sidebar";
import { Input } from "../ui/input";
import Image from "next/image";
import { getServerSession } from "next-auth";
import { options } from "@/app/api/auth/[...nextauth]/options";
import { BackendURL } from "@/config/env";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import LogoutButton from "./LogoutButton";
import Link from "next/link";

async function GetBalance(id: string) {
  try {
    const res = await fetch(
      `${BackendURL}/api/v1/data/balance?affiliateId=${id}`
    );

    if (res.status !== 200) {
      return 0;
    }

    const data = await res.json();

    return data.data || 0;
  } catch (error) {
    console.log(error);

    return 0;
  }
}

export default async function Navbar() {
  const session = await getServerSession(options);

  const balance = await GetBalance(session?.user.affiliateId || "");

  return (
    <header className="flex z-30 w-full sm:h-16 h-14 shrink-0 items-center gap-2 px-4 max-sm:border-b sm:bg-[#F8F9FA] bg-white sticky top-0">
      <SidebarTrigger className="-ml-1 md:hidden flex" />

      <div className="flex justify-between w-full">
        <div className="flex sm:opacity-100 opacity-0 flex-col text-sm">
          <div className="text-xs">Hi Good Morning,</div>
          <div className="font-semibold">{session?.user.name}</div>
        </div>

        <div className="flex items-center sm:gap-4 gap-2">
          <div className="sm:flex hidden bg-[#152C28] sm:rounded-2xl rounded-lg sm:pl-5 pl-2 sm:h-10 h-8">
            <div className="py-1 sm:pr-5 pr-2 flex sm:gap-2 gap-1 items-center">
              <Image
                src={"/images/wallet.svg"}
                alt={"wallet"}
                width={20}
                height={20}
                className="sm:hidden flex"
              />

              <div className="flex flex-col gap-0">
                <div className="text-[#F5F5F5] text-xs sm:flex hidden relative top-1">
                  Balance
                </div>
                <div className="text-[#20F6CA] sm:text-sm text-xs relative max-sm:top-px font-semibold leading-none">
                  ${balance}
                </div>
              </div>
            </div>

            <div className="bg-walletGradient rounded-2xl px-6 gap-2 sm:flex hidden items-center">
              <Image
                src={"/images/wallet.svg"}
                alt={"wallet"}
                width={25}
                height={25}
              />

              <span className="text-white relative top-px font-semibold text-sm">
                Wallet
              </span>
            </div>
          </div>

          <div className="relative sm:flex hidden">
            <SearchIcon className="absolute top-1/2 -translate-y-1/2 left-2" />
            <Input
              className="w-fit h-10 bg-white pl-10"
              placeholder="Search..."
            />
          </div>

          <div className="flex items-center gap-2">
            <DropdownMenu>
              <DropdownMenuTrigger asChild>
                <div className="bg-gray p-1 rounded-full cursor-pointer">
                  <User2 />
                </div>
              </DropdownMenuTrigger>
              <DropdownMenuContent>
                <Link href={"/settings"} className="text-left cursor-pointer">
                  <DropdownMenuItem className="cursor-pointer">
                    Settings
                  </DropdownMenuItem>
                </Link>

                <LogoutButton />
              </DropdownMenuContent>
            </DropdownMenu>

            <div className="text-xs sm:flex flex-col hidden">
              <div>{session?.user.name}</div>
              <div className="text-gray text-xs">
                #{session?.user.affiliateId}
              </div>
            </div>
            <div className="sm:hidden flex flex-col gap-0">
              <div className="text-[#686868] text-xs relative top-px">
                Balance
              </div>
              <div className="text-[#015C5D] sm:text-sm text-xs relative max-sm:top-px font-semibold leading-none">
                ${balance}
              </div>
            </div>
          </div>
        </div>
      </div>
    </header>
  );
}
