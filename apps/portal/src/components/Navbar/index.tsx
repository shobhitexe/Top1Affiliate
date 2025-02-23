import { SearchIcon, User2 } from "lucide-react";
import { SidebarTrigger } from "../ui/sidebar";
import { Input } from "../ui/input";

export default function Navbar() {
  return (
    <header className="flex z-30 w-full sm:h-16 h-14 shrink-0 items-center gap-2 px-4 max-sm:border-b sm:bg-[#F8F9FA] bg-white sticky top-0">
      <SidebarTrigger className="-ml-1 md:hidden flex" />

      <div className="flex justify-between w-full">
        <div className="flex sm:opacity-100 opacity-0 flex-col text-sm">
          <div className="text-xs">Hi Good Morning,</div>
          <div className="font-semibold">Kilian</div>
        </div>

        <div className="flex items-center gap-4">
          <div className="relative sm:flex hidden">
            <SearchIcon className="absolute top-1/2 -translate-y-1/2 left-2" />
            <Input
              className="w-fit h-10 bg-white pl-10"
              placeholder="Search..."
            />
          </div>

          <div className="flex items-center gap-2">
            {" "}
            <div className="bg-gray p-1 rounded-full">
              <User2 />
            </div>
            <div className="text-xs">
              <div>Killian</div>
              <div className="text-gray text-xs">#32648723</div>
            </div>
          </div>
        </div>
      </div>
    </header>
  );
}
