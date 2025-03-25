"use server";

import { BackendURL } from "@/config/env";
import { revalidatePath } from "next/cache";

export async function SwitchStatus(id: string) {
  try {
    const res = await fetch(`${BackendURL}/api/v1/admin/block`, {
      method: "POST",
      body: JSON.stringify({ id }),
    });

    if (res.status !== 200) {
      return false;
    }

    return true;
  } catch (error) {
    console.log(error);
    return false;
  } finally {
    revalidatePath("/list");
  }
}
