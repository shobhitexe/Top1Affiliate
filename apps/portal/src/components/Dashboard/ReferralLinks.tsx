"use client";

import { useToast } from "@/hooks/use-toast";
import { Button } from "../ui/button";
import { useSession } from "next-auth/react";

export default function ReferralLinks() {
  const session = useSession();

  const Links = [
    {
      title: "Client Referral Link",
      link: `${session.data?.user.Clientlink}`,
      description:
        "Share your referral link by copying and sending it to your Clients or sharing it on Social media.",
    },
    {
      title: "Affiliate Referral Link",
      link: `${session.data?.user.Sublink}`,
      description:
        "You can also share your referral link by copying and sending it to your Affilliates or sharing it on Social media.",
    },
  ];
  return (
    <div className="bg-white shadow-sm rounded-2xl p-4 flex flex-col gap-6">
      {Links.map((item) => (
        <ReferralLinkComponent key={item.title} {...item} />
      ))}
    </div>
  );
}

function ReferralLinkComponent({
  title,
  link,
  description,
}: {
  title: string;
  link: string;
  description: string;
}) {
  const { toast } = useToast();

  async function copyToClipboard() {
    await navigator.clipboard.writeText(link);

    toast({ title: "Copied to clipboard", variant: "default" });
  }

  return (
    <div className="flex flex-col gap-4">
      <div>
        <div className="font-semibold tracking-wide text-sm">{title}</div>
        <div className="text-[#6A7179] text-sm">{description}</div>
      </div>

      <div className="flex lg:flex-row flex-col lg:items-center gap-2">
        <div className="text-[#6A7179] border border-[#E2E8F0] rounded-2xl p-3 w-full text-sm">
          {link}
        </div>
        <Button variant={"outline"} className="h-12" onClick={copyToClipboard}>
          Copy Link
        </Button>
      </div>
    </div>
  );
}
