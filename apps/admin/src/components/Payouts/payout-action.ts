"use server";

import { BackendURL } from "@/config/env";
import { revalidatePath } from "next/cache";

export async function DeclinePayoutAction(id: string) {
  try {
    const res = await fetch(
      `${BackendURL}/api/v1/admin/wallet/payout/decline`,
      {
        method: "POST",
        body: JSON.stringify({ id }),
      }
    );

    if (res.status !== 200) {
      return false;
    }

    revalidatePath("/payouts/pending");
    return true;
  } catch (error) {
    console.log(error);

    return false;
  }
}

export async function AcceptPayoutAction(id: string, amount: number) {
  try {
    const res = await fetch(`${BackendURL}/api/v1/admin/wallet/payout/accept`, {
      method: "POST",
      body: JSON.stringify({ id, amount: Number(amount) }),
    });

    if (res.status !== 200) {
      return false;
    }

    revalidatePath("/payouts/pending");
    return true;
  } catch (error) {
    console.log(error);

    return false;
  }
}
