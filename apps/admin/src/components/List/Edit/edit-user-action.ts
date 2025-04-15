"use server";

import { BackendURL } from "@/config/env";
import { revalidatePath } from "next/cache";

export async function EditUserAction(
  name: string,
  country: string,
  commission: number,
  link: string,
  id: string
) {
  try {
    const res = await fetch(`${BackendURL}/api/v1/admin/edit`, {
      method: "POST",
      body: JSON.stringify({ name, country, commission, id, link }),
    });

    if (res.status !== 200) {
      return false;
    }

    revalidatePath(`/list/edit/${id}`);

    return true;
  } catch (error) {
    console.log(error);

    return false;
  }
}
