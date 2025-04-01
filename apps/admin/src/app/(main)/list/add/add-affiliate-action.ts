"use server";

import { BackendURL } from "@/config/env";

type Affiliate = {
  name: string;
  affiliateId: string;
  country: string;
  commission: number;
  password: string;
};

export async function AddNewAffiliateAction(data: Affiliate) {
  try {
    const res = await fetch(`${BackendURL}/api/v1/admin/affiliate`, {
      method: "POST",
      body: JSON.stringify(data),
    });

    if (res.status !== 200) {
      return false;
    }

    return true;
  } catch (error) {
    console.log(error);

    return false;
  }
}
