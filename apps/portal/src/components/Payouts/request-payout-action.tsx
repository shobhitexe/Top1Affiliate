"use server";

import { BackendURL } from "@/config/env";

export async function RequestPayoutAction(
  amount: number,
  id: string,
  type: string
) {
  try {
    const res = await fetch(`${BackendURL}/api/v1/wallet/payout`, {
      method: "POST",
      body: JSON.stringify({ amount, id, type }),
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
