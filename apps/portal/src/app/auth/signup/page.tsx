import { RegisterForm } from "@/components";
import Image from "next/image";

export default function RegisterPage() {
  return (
    <div className="grid min-h-svh lg:grid-cols-2">
      <div className="relative hidden bg-muted lg:block">
        <Image
          src="/images/auth/bg.svg"
          alt="Image"
          className="absolute inset-0 h-full w-full object-cover dark:brightness-[0.2] dark:grayscale"
          width={1079}
          height={1172}
        />
      </div>
      <div className="flex flex-col gap-4 p-6 md:p-10">
        <div className="flex flex-1 items-center justify-center">
          <div className="w-full max-w-sm">
            <RegisterForm />
          </div>
        </div>
      </div>
    </div>
  );
}
