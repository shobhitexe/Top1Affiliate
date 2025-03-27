"use server";

import { BackendURL } from "@/config/env";
import { revalidatePath } from "next/cache";

export async function UpdateWalletDetails(data: FormData, id: string) {
  const formData = {
    iban: data.get("iban"),
    swift: data.get("swift"),
    bankName: data.get("bankName"),
    chainName: data.get("chainName"),
    walletAddress: data.get("walletAddress"),
  };

  try {
    const res = await fetch(`${BackendURL}/api/v1/wallet/details`, {
      method: "POST",
      body: JSON.stringify({ ...formData, id }),
    });

    if (res.status !== 200) {
      return false;
    }

    revalidatePath("/settings");

    return true;
  } catch (error) {
    console.log(error);
    return false;
  }
}
