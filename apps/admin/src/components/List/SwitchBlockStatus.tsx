"use client";

import { useToast } from "@/hooks/use-toast";
import { Switch } from "../ui/switch";
import { SwitchStatus } from "./switch-status-action";

export default function SwitchBlockStatus({
  status,
  id,
}: {
  status: boolean;
  id: string;
}) {
  const { toast } = useToast();

  async function ChangeStatusClient() {
    try {
      const res = await SwitchStatus(id);

      if (res === true) {
        toast({ title: "Changed Status" });
        return;
      }

      toast({ title: "Failed to Changed Status", variant: "destructive" });
    } catch (error) {
      console.log(error);
      toast({ title: "Failed to Changed Status", variant: "destructive" });
    }
  }

  return (
    <div className="flex items-center gap-2">
      <Switch defaultChecked={status} onCheckedChange={ChangeStatusClient} />
      <div>{status ? "Blocked" : ""}</div>
    </div>
  );
}
