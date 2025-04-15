import { getServerSession } from "next-auth";
import { Button } from "../ui/button";
import { options } from "@/app/api/auth/[...nextauth]/options";

export default async function ReferralLinks() {
  const session = await getServerSession(options);

  const Links = [
    {
      title: "Client Referral Link",
      link: `${session?.user.link}`,
      description:
        "Share your referral link by copying and sending it to your Clients or sharing it on Social media.",
    },
    {
      title: "Affiliate Referral Link",
      link: `${session?.user.link}`,
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
        <Button variant={"outline"} className="h-12">
          Copy Link
        </Button>
      </div>
    </div>
  );
}
