import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Button } from "../ui/button";
import { useToast } from "@/hooks/use-toast";
import { AcceptPayoutAction } from "./payout-action";

export default function AcceptPayout({
  id,
  amount,

  method,

  iban,
  swiftCode,
  bankName,

  chainName,
  walletAddress,
}: {
  id: string;
  amount: number;

  method: string;

  iban: string;
  swiftCode: string;
  bankName: string;

  chainName: string;
  walletAddress: string;
}) {
  const { toast } = useToast();

  async function AcceptPayoutClient() {
    try {
      const res = await AcceptPayoutAction(id, amount);

      if (res === true) {
        toast({ title: "Approved" });
        return;
      }

      toast({ title: "Failed", variant: "destructive" });
    } catch (error) {
      toast({ title: "Failed", variant: "destructive" });
      console.log(error);
    }
  }

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button>Approve</Button>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Are you absolutely sure?</DialogTitle>
          <div className="pt-5" />
          <div>Approve payout of ${amount}</div>

          <div className="pt-5" />
          {method === "bank" ? (
            <div>
              <div>IBAN: {iban}</div>
              <div>SWIFT CODE: {swiftCode}</div>
              <div>BANK NAME: {bankName}</div>
            </div>
          ) : (
            <div>
              <div>CHAIN NAME: {chainName}</div>
              <div>WALLET ADDRESS: {walletAddress}</div>
            </div>
          )}

          <div className="pt-5" />
          <Button onClick={AcceptPayoutClient} className="mt-5">
            Approve
          </Button>
        </DialogHeader>
      </DialogContent>
    </Dialog>
  );
}
