import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Button } from "../ui/button";
import { useToast } from "@/hooks/use-toast";
import { DeclinePayoutAction } from "./payout-action";

export default function DeclinePayout({ id }: { id: string }) {
  const { toast } = useToast();

  async function DeclinePayoutClient() {
    try {
      const res = await DeclinePayoutAction(id);

      if (res === true) {
        toast({ title: "Declined" });
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
        <Button variant={"destructive"}>Decline</Button>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Are you absolutely sure?</DialogTitle>
          <div className="pt-5" />
          <Button
            variant={"destructive"}
            onClick={DeclinePayoutClient}
            className="mt-5"
          >
            Decline
          </Button>
        </DialogHeader>
      </DialogContent>
    </Dialog>
  );
}
