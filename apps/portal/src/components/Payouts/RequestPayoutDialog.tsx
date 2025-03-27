"use client";

import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Button } from "../ui/button";
import PayoutsIcon from "../Sidebar/Icons/payouts";
import { Input } from "../ui/input";
import { useToast } from "@/hooks/use-toast";
import { useState } from "react";
import { useSession } from "next-auth/react";
import { RequestPayoutAction } from "./request-payout-action";
import Image from "next/image";

export default function RequestPayoutDialog({
  balance,
  type,
}: {
  balance: number;
  type: "payout" | "transfer";
}) {
  const { toast } = useToast();

  const session = useSession();

  const [amount, setAmount] = useState(0);

  async function RequestPayoutClient() {
    try {
      const res = await RequestPayoutAction(
        amount,
        session.data?.user.id || "",
        type
      );

      if (res === true) {
        toast({ title: "Request submitted" });
        setAmount(0);
        return;
      }

      toast({ title: "Failed to request payout", variant: "destructive" });
    } catch (error) {
      console.log(error);
      toast({ title: "Failed to request payout", variant: "destructive" });
    }
  }

  return (
    <Dialog>
      <DialogTrigger asChild>
        {type === "payout" ? (
          <Button size={"lg"} className="font-semibold rounded-xl">
            <PayoutsIcon fill="white" />{" "}
            <span className="relative top-px">Request Payout</span>
          </Button>
        ) : (
          <Button size={"lg"} className="font-semibold bg-[#237C81] rounded-xl">
            <Image
              src={"/images/transfer.svg"}
              alt={"transfer"}
              width={22}
              height={20}
            />{" "}
            <span className="relative top-px">Transfer to Trading Acc</span>
          </Button>
        )}
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Request Payout</DialogTitle>
          <form
            className="flex flex-col gap-4 pt-4"
            action={RequestPayoutClient}
          >
            <Input
              name="amount"
              type="number"
              id="amount"
              placeholder="Enter Amount"
              required
              onChange={(e) => setAmount(Number(e.target.value))}
              max={balance}
            />

            <Button>Request</Button>
          </form>
        </DialogHeader>
      </DialogContent>
    </Dialog>
  );
}
