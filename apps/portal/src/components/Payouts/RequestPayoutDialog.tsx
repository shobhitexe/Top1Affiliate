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

import useSWR from "swr";
import fetcher from "@/lib/fetcher";
import { BackendURL } from "@/config/env";
import { WalletDetails } from "@/types";
import Link from "next/link";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

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
  const [method, setMethod] = useState("");

  const { data } = useSWR<{ data: WalletDetails }>(
    `${BackendURL}/api/v1/wallet/details?id=${session.data?.user.id || ""}`,
    fetcher
  );

  async function RequestPayoutClient() {
    try {
      const res = await RequestPayoutAction(
        amount,
        session.data?.user.id || "",
        type,
        method
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
          {!data?.data ? (
            <div className="mt-4">
              Please Update Wallet details before requesting payout in settings
              page <br />
              <Link
                href={"/settings"}
                className="underline cursor-pointer text-[#237C81]"
              >
                Take me there
              </Link>
            </div>
          ) : (
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

              <Select
                required
                value={method}
                onValueChange={(e) => setMethod(e)}
              >
                <SelectTrigger className="">
                  <SelectValue placeholder="Wallet" />
                </SelectTrigger>
                <SelectContent>
                  {data?.data.iban.length > 3 && (
                    <SelectItem value="bank">Bank</SelectItem>
                  )}
                  {data.data.walletAddress.length > 5 && (
                    <SelectItem value="crypto">Crypto</SelectItem>
                  )}
                </SelectContent>
              </Select>

              <Button>Request</Button>
            </form>
          )}
        </DialogHeader>
      </DialogContent>
    </Dialog>
  );
}
