"use client";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useEffect, useState } from "react";
import { AddNewAffiliateAction } from "./add-affiliate-action";
import { useToast } from "@/hooks/use-toast";
import { useSearchParams } from "next/navigation";

// type Affiliate = {
//   id: string;
//   name: string;
//   affiliateId: "string";
//   country: string;
//   commission: number;
// };

export default function AddAffiliate() {
  const { toast } = useToast();

  const searchparams = useSearchParams();

  const params = {
    name: searchparams.get("name"),
    id: searchparams.get("id"),
  };

  const [data, setData] = useState({
    addedBy: Number(params.id) || 0,
    id: "",
    name: "",
    affiliateId: "",
    country: "",
    password: "",
    commission: 0,
    Clientlink: "",
    Sublink: "",
  });

  useEffect(() => {
    setData((prev) => ({ ...prev, addedBy: Number(params.id) || 0 }));
  }, [params.id]);

  async function SubmitActionClient() {
    try {
      const res = await AddNewAffiliateAction(data);

      if (res !== true) {
        toast({ title: "Failed", variant: "destructive" });
        return;
      }

      toast({ title: "New Affiliate registered" });

      setData({
        addedBy: Number(params.id) || 0,
        id: "",
        name: "",
        affiliateId: "",
        country: "",
        password: "",
        commission: 0,
        Clientlink: "",
        Sublink: "",
      });
    } catch (error) {
      console.log(error);
      toast({ title: "Failed", variant: "destructive" });
    }
  }

  return (
    <form
      className="flex flex-col sm:gap-4 gap-2 bg-white sm:p-5 p-3 shadow-sm rounded-2xl"
      action={SubmitActionClient}
    >
      <div className="font-semibold text-lg">
        Add new Affiliate {params.name && `under (${params.name})`}
      </div>

      <div className="grid sm:grid-cols-2 gap-3">
        <div className="grid w-full items-center gap-1.5">
          <Label htmlFor="affiliateid">Affiliate ID</Label>
          <Input
            name="affiliateid"
            type="text"
            id="affiliateid"
            placeholder="Affiliate ID"
            required
            value={data.affiliateId}
            onChange={(e) =>
              setData((prev) => ({ ...prev, affiliateId: e.target.value }))
            }
          />
        </div>

        <div className="grid w-full items-center gap-1.5">
          <Label htmlFor="name">Name</Label>
          <Input
            name="name"
            type="text"
            id="name"
            placeholder="Name"
            required
            value={data.name}
            onChange={(e) =>
              setData((prev) => ({ ...prev, name: e.target.value }))
            }
          />
        </div>
      </div>

      <div className="grid sm:grid-cols-2 gap-3">
        <div className="grid w-full items-center gap-1.5">
          <Label htmlFor="country">Country</Label>
          <Input
            name="country"
            type="text"
            id="country"
            placeholder="country"
            required
            value={data.country}
            onChange={(e) =>
              setData((prev) => ({ ...prev, country: e.target.value }))
            }
          />
        </div>

        <div className="grid w-full items-center gap-1.5">
          <Label htmlFor="name">Commission</Label>
          <Input
            name="commission"
            type="number"
            id="commission"
            placeholder="commission"
            required
            value={data.commission}
            onChange={(e) =>
              setData((prev) => ({
                ...prev,
                commission: Number(e.target.value),
              }))
            }
          />
        </div>
      </div>

      <div className="grid sm:grid-cols-2 gap-3">
        <div className="grid w-full items-center gap-1.5">
          <Label htmlFor="link">Client Affiliate Link</Label>
          <Input
            name="link"
            type="text"
            id="link"
            placeholder="Client Affiliate Link"
            required
            value={data.Clientlink}
            onChange={(e) =>
              setData((prev) => ({
                ...prev,
                Clientlink: e.target.value,
              }))
            }
          />
        </div>
        <div className="grid w-full items-center gap-1.5">
          <Label htmlFor="link">Sub Affiliate Link</Label>
          <Input
            name="link"
            type="text"
            id="link"
            placeholder="Sub Affiliate Link"
            required
            value={data.Sublink}
            onChange={(e) =>
              setData((prev) => ({
                ...prev,
                Sublink: e.target.value,
              }))
            }
          />
        </div>
      </div>

      <div className="grid w-full items-center gap-1.5">
        <Label htmlFor="password">Password</Label>
        <Input
          name="password"
          type="text"
          id="password"
          placeholder="password"
          required
          value={data.password}
          onChange={(e) =>
            setData((prev) => ({
              ...prev,
              password: e.target.value,
            }))
          }
        />
      </div>

      <Button size={"lg"} className="mt-5 sm:w-fit">
        Create
      </Button>
    </form>
  );
}
