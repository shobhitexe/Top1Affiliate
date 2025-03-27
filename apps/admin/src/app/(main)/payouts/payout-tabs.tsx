"use client";

import { buttonVariants } from "@/components/ui/button";
import Link from "next/link";
import { usePathname } from "next/navigation";

const Tabs = [
  { title: "Pending", link: "/payouts/pending" },
  { title: "Paid", link: "/payouts/paid" },
  { title: "Rejected", link: "/payouts/rejected" },
];

export default function PayoutTabs() {
  const pathname = usePathname();

  console.log(pathname);

  return (
    <div className="flex items-center gap-2">
      {Tabs.map((item) => (
        <Link
          key={item.title}
          href={item.link}
          className={`${buttonVariants({
            variant: item.link === pathname ? "default" : "outline",
          })}`}
        >
          {item.title}
        </Link>
      ))}
    </div>
  );
}
