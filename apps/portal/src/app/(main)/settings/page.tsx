import { options } from "@/app/api/auth/[...nextauth]/options";
import { AddUpdateWalletDetails } from "@/components";
import { Card, CardContent, CardHeader } from "@/components/ui/card";
import { BackendURL } from "@/config/env";
import { WalletDetails } from "@/types";
import { Bitcoin, CreditCard, User } from "lucide-react";
import { getServerSession } from "next-auth";

async function GetWalletDetails(id: string) {
  try {
    const res = await fetch(`${BackendURL}/api/v1/wallet/details?id=${id}`);

    if (res.status !== 200) {
      return {
        iban: "N/A",
        swift: "N/A",
        bankName: "N/A",
        chainName: "N/A",
        walletAddress: "N/A",
      };
    }

    const data = await res.json();

    return (
      data.data || {
        iban: "N/A",
        swift: "N/A",
        bankName: "N/A",
        chainName: "N/A",
        walletAddress: "N/A",
      }
    );
  } catch (error) {
    console.log(error);

    return {
      iban: "N/A",
      swift: "N/A",
      bankName: "N/A",
      chainName: "N/A",
      walletAddress: "N/A",
    };
  }
}

export default async function page() {
  const session = await getServerSession(options);

  const wallet: WalletDetails = await GetWalletDetails(session?.user.id || "");

  return (
    <div className="flex flex-col items-center justify-center bg-white p-2 shadow-sm rounded-2xl">
      <Card className="w-full border-0">
        <CardHeader className="pb-2 sm:px-4 px-2">
          <div>
            <h2 className="font-semibold text-lg">Profile</h2>
            <p className="text-sm text-muted-foreground">Account information</p>
          </div>
        </CardHeader>
        <CardContent className="sm:px-4 px-2">
          <div className="space-y-4">
            <div className="flex items-center space-x-3 rounded-md border p-3">
              <User className="h-5 w-5 text-primary" />
              <div className="flex-1 space-y-1">
                <p className="text-sm font-medium leading-none">Name</p>
                <p className="text-sm text-muted-foreground">
                  {session?.user.name}
                </p>
              </div>
            </div>

            <div className="flex items-center space-x-3 rounded-md border p-3">
              <CreditCard className="h-5 w-5 text-primary" />
              <div className="flex-1 space-y-1">
                <p className="text-sm font-medium leading-none">Affiliate ID</p>
                <p className="text-sm text-muted-foreground">
                  {session?.user.affiliateId}
                </p>
              </div>
            </div>
          </div>
        </CardContent>
      </Card>

      <Card className="w-full border-0">
        <CardHeader className="pb-2 p-0 py-4 sm:px-4 px-2">
          <div>
            <h2 className="font-semibold text-lg">Wallet</h2>
            <p className="text-sm text-muted-foreground">Wallet information</p>
          </div>
        </CardHeader>
        <CardContent className="sm:px-4 px-2">
          <div className="space-y-4">
            <div className="flex flex-col gap-4 space-x-3 rounded-md border">
              <div className="flex items-center gap-2 border-b p-3 justify-between">
                <div className="flex items-center gap-2">
                  <div>Banking Details</div>
                  <CreditCard className="h-5 w-5 text-primary relative -top-1" />
                </div>

                <AddUpdateWalletDetails type="bank" wallet={wallet} />
              </div>

              <div className="flex flex-col gap-4 justify-between w-full pb-2">
                <div className="flex-1 space-y-1">
                  <p className="text-sm font-medium leading-none">
                    IBAN Number
                  </p>
                  <p className="text-sm text-muted-foreground">{wallet.iban}</p>
                </div>

                <div className="flex-1 space-y-1">
                  <p className="text-sm font-medium leading-none">SWIFT Code</p>
                  <p className="text-sm text-muted-foreground">
                    {wallet.swift}
                  </p>
                </div>

                <div className="flex-1 space-y-1">
                  <p className="text-sm font-medium leading-none">Bank Name</p>
                  <p className="text-sm text-muted-foreground">
                    {wallet.bankName}
                  </p>
                </div>
              </div>
            </div>

            <div className="flex flex-col gap-4 space-x-3 rounded-md border">
              <div className="flex items-center gap-2 border-b p-3 justify-between">
                <div className="flex items-center gap-2">
                  <div>Crypto Details</div>
                  <Bitcoin className="h-5 w-5 text-primary relative -top-1" />
                </div>

                <AddUpdateWalletDetails type="crypto" wallet={wallet} />
              </div>

              <div className="flex flex-col gap-4 justify-between w-full pb-2">
                <div className="flex-1 space-y-1">
                  <p className="text-sm font-medium leading-none">Chain Name</p>
                  <p className="text-sm text-muted-foreground">
                    {wallet.chainName}
                  </p>
                </div>

                <div className="flex-1 space-y-1">
                  <p className="text-sm font-medium leading-none">
                    Wallet Address
                  </p>
                  <p className="text-sm text-muted-foreground">
                    {wallet.walletAddress}
                  </p>
                </div>
              </div>
            </div>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}
