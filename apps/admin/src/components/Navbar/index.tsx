import { User2 } from "lucide-react";
import { SidebarTrigger } from "../ui/sidebar";
// import { Input } from "../ui/input";
import { getServerSession } from "next-auth";
import { options } from "@/app/api/auth/[...nextauth]/options";

export default async function Navbar() {
  const session = await getServerSession(options);

  return (
    <header className="flex z-30 w-full sm:h-16 h-14 shrink-0 items-center gap-2 px-4 max-sm:border-b sm:bg-[#F8F9FA] bg-white sticky top-0">
      <SidebarTrigger className="-ml-1 md:hidden flex" />

      <div className="flex justify-between w-full">
        <div className="flex sm:opacity-100 opacity-0 flex-col text-sm">
          <div className="text-xs">Hi Good Morning,</div>
          <div className="font-semibold">{session?.user.username}</div>
        </div>

        <div className="flex items-center sm:gap-4 gap-2">
          {/* <div className="relative sm:flex hidden">
            <SearchIcon className="absolute top-1/2 -translate-y-1/2 left-2" />
            <Input
              className="w-fit h-10 bg-white pl-10"
              placeholder="Search..."
            />
          </div> */}

          <div className="flex items-center gap-2">
            {" "}
            <div className="bg-gray p-1 rounded-full">
              <User2 />
            </div>
            <div className="text-xs sm:flex flex-col hidden">
              <div>{session?.user.name}</div>
              <div className="text-gray text-xs">
                {/* #{session?.user.affiliateId} */}
              </div>
            </div>
            <div className="sm:hidden flex flex-col gap-0">
              <div className="text-[#686868] text-xs relative top-px">
                {session?.user.username}
              </div>
            </div>
          </div>
        </div>
      </div>
    </header>
  );
}
