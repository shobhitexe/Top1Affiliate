"use client";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useToast } from "@/hooks/use-toast";
import { Affiliate } from "@/types";
import { EditUserAction } from "./edit-user-action";

export default function EditUser({
  affiliate,
  id,
}: {
  affiliate: Affiliate;
  id: string;
}) {
  const { toast } = useToast();

  async function EditAdminClient(formData: FormData) {
    const updatedAffiliate = {
      name: formData.get("name") || "",
      country: formData.get("country") || "",
      commission: formData.get("commission") || "",
      Clientlink: formData.get("Clientlink") || "",
      Sublink: formData.get("Sublink") || "",
      balance: formData.get("balance") || 0,
      password: formData.get("password") || "",
    };

    try {
      const res = await EditUserAction(
        updatedAffiliate.name as string,
        updatedAffiliate.country as string,
        Number(updatedAffiliate.commission),
        updatedAffiliate.Clientlink as string,
        updatedAffiliate.Sublink as string,
        updatedAffiliate.balance as number,
        updatedAffiliate.password as string,
        id
      );

      if (res === true) {
        toast({ title: "Edited" });
        return;
      }
      toast({ title: "Failed to Edit", variant: "destructive" });
    } catch (error) {
      toast({ title: "Failed to Edit", variant: "destructive" });
      console.log(error);
    }
  }

  return (
    <form
      action={EditAdminClient}
      className="flex flex-col sm:gap-4 gap-2 bg-white p-5 shadow-sm rounded-2xl"
    >
      <div className="font-semibold text-lg">Edit {affiliate.name}</div>

      <div className="grid sm:grid-cols-2 gap-3">
        <div className="grid w-full items-center gap-1.5">
          <Label htmlFor="affiliateid">Affiliate ID</Label>
          <Input
            type="text"
            id="affiliateid"
            name="affiliateid"
            placeholder="Affiliate ID"
            defaultValue={affiliate.affiliateId}
            disabled
          />
        </div>

        <div className="grid w-full items-center gap-1.5">
          <Label htmlFor="name">Name</Label>
          <Input
            type="text"
            id="name"
            name="name"
            placeholder="Name"
            defaultValue={affiliate.name}
          />
        </div>
      </div>

      <div className="grid sm:grid-cols-2 gap-3">
        <div className="grid w-full items-center gap-1.5">
          <Label htmlFor="country">Country</Label>
          <Input
            type="text"
            id="country"
            name="country"
            placeholder="country"
            defaultValue={affiliate.country}
          />
        </div>

        <div className="grid w-full items-center gap-1.5">
          <Label htmlFor="commission">Commission</Label>
          <Input
            type="number"
            id="commission"
            name="commission"
            placeholder="commission"
            defaultValue={affiliate.commission}
          />
        </div>
      </div>

      <div className="grid sm:grid-cols-2 gap-3">
        <div className="grid w-full items-center gap-1.5">
          <Label htmlFor="Clientlink">Client Affiliate Link</Label>
          <Input
            type="text"
            id="Clientlink"
            name="Clientlink"
            placeholder="client link"
            defaultValue={affiliate.Clientlink}
          />
        </div>

        <div className="grid w-full items-center gap-1.5">
          <Label htmlFor="Sublink">Sub Affiliate Link</Label>
          <Input
            type="text"
            id="Sublink"
            name="Sublink"
            placeholder="sub link"
            defaultValue={affiliate.Sublink}
          />
        </div>
      </div>

      <div className="grid w-full items-center gap-1.5">
        <Label htmlFor="balance">Balance</Label>
        <Input
          type="number"
          id="balance"
          name="balance"
          placeholder="Balance"
          step={0.01}
          defaultValue={affiliate.balance}
        />
      </div>

      <div className="grid w-full items-center gap-1.5">
        <Label htmlFor="password">Password</Label>
        <Input
          type="text"
          id="password"
          name="password"
          placeholder="Password"
        />
      </div>

      <Button size={"lg"} className="mt-5 sm:w-fit">
        Edit
      </Button>
    </form>
  );
}
