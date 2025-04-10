"use client";

import Link from "next/link";
import { buttonVariants } from "../ui/button";
import { usePathname } from "next/navigation";
import { useSession } from "next-auth/react";

export default function ViewButtons() {
  const pathname = usePathname();

  const session = useSession();

  const Buttons = [
    { name: "Tree", link: `/sub-affiliates/${session.data?.user.id}/tree` },
    { name: "List", link: `/sub-affiliates/${session.data?.user.id}/list` },
    { name: "Table", link: `/sub-affiliates/${session.data?.user.id}/table` },
  ];

  return (
    <div className="flex items-center justify-between">
      <div className="flex gap-2">
        {Buttons.map((button) => (
          <Link
            key={button.name}
            href={button.link}
            className={`${buttonVariants({
              variant: button.link === pathname ? "default" : "outline",
            })}`}
          >
            {button.name}
          </Link>
        ))}
      </div>
    </div>
  );
}
