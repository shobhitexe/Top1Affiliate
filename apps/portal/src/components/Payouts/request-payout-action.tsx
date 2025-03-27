"use server";

import { BackendURL } from "@/config/env";
import { revalidatePath } from "next/cache";

export async function RequestPayoutAction(
  amount: number,
  id: string,
  type: string,
  method: string
) {
  try {
    const res = await fetch(`${BackendURL}/api/v1/wallet/payout`, {
      method: "POST",
      body: JSON.stringify({ amount, id, type, method }),
    });

    if (res.status !== 200) {
      return false;
    }

    revalidatePath("/payouts");
    return true;
  } catch (error) {
    console.log(error);
    return false;
  }
}
