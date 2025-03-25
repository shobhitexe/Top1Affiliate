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
    };

    try {
      const res = await EditUserAction(
        updatedAffiliate.name as string,
        updatedAffiliate.country as string,
        Number(updatedAffiliate.commission),
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

      <Button size={"lg"} className="mt-5 sm:w-fit">
        Edit
      </Button>
    </form>
  );
}
