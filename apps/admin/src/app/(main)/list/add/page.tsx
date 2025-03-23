import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";

const affiliate = {
  id: "",
  name: "",
  affiliateId: "",
  country: "",
  commission: 0,
};

// type Affiliate = {
//   id: string;
//   name: string;
//   affiliateId: "string";
//   country: string;
//   commission: number;
// };

export default async function page() {
  return (
    <div className="flex flex-col sm:gap-4 gap-2 bg-white p-5 shadow-sm rounded-2xl">
      <div className="font-semibold text-lg">Add new Affiliate</div>

      <div className="grid sm:grid-cols-2 gap-3">
        <div className="grid w-full items-center gap-1.5">
          <Label htmlFor="affiliateid">Affiliate ID</Label>
          <Input
            type="text"
            id="affiliateid"
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
            placeholder="country"
            defaultValue={affiliate.country}
          />
        </div>

        <div className="grid w-full items-center gap-1.5">
          <Label htmlFor="name">Commission</Label>
          <Input
            type="number"
            id="commission"
            placeholder="commission"
            defaultValue={affiliate.commission}
          />
        </div>
      </div>

      <Button size={"lg"} className="mt-5 sm:w-fit">
        Create
      </Button>
    </div>
  );
}
