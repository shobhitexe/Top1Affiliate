"use client";

import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Button } from "../ui/button";
import { WalletDetails } from "@/types";
import { Input } from "../ui/input";
import { UpdateWalletDetails } from "./update-wallet-action";
import { useSession } from "next-auth/react";
import { useToast } from "@/hooks/use-toast";

export default function AddUpdateWalletDetails({
  type,
  wallet,
}: {
  type: "bank" | "crypto";
  wallet: WalletDetails;
}) {
  const { toast } = useToast();

  const session = useSession();

  async function UpdateDetailsClient(data: FormData) {
    try {
      const res = await UpdateWalletDetails(data, session.data?.user.id || "");

      if (res === true) {
        toast({ title: "Details Updated" });
        return;
      }

      toast({ title: "Failed", variant: "destructive" });
    } catch (error) {
      console.log(error);
      toast({ title: "Failed", variant: "destructive" });
    }
  }

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button>Update</Button>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Update {type} details</DialogTitle>
          <form
            className="pt-4 flex flex-col gap-4"
            action={UpdateDetailsClient}
          >
            <div
              className={`${
                type === "bank" ? "" : "hidden"
              } flex flex-col gap-4`}
            >
              <Input
                type="text"
                id="iban"
                name="iban"
                placeholder="IBAN Number"
                defaultValue={wallet.iban === "N/A" ? "" : wallet.iban}
              />

              <Input
                type="text"
                id="swift"
                name="swift"
                placeholder="Swift Code"
                defaultValue={wallet.swift === "N/A" ? "" : wallet.swift}
              />

              <Input
                type="text"
                id="bankName"
                name="bankName"
                placeholder="Bank Name"
                defaultValue={wallet.bankName === "N/A" ? "" : wallet.bankName}
              />
            </div>
            <div
              className={`${
                type === "bank" ? "hidden" : ""
              } flex flex-col gap-4`}
            >
              <Input
                type="text"
                id="chainName"
                name="chainName"
                placeholder="Chain Name"
                defaultValue={
                  wallet.chainName === "N/A" ? "" : wallet.chainName
                }
              />

              <Input
                type="text"
                id="walletAddress"
                name="walletAddress"
                placeholder="Wallet Address"
                defaultValue={
                  wallet.walletAddress === "N/A" ? "" : wallet.walletAddress
                }
              />
            </div>

            <Button>Update</Button>
          </form>
        </DialogHeader>
      </DialogContent>
    </Dialog>
  );
}
