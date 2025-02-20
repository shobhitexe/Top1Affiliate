import { cn } from "@/lib/utils";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import Link from "next/link";

export function RegisterForm({
  className,
  ...props
}: React.ComponentPropsWithoutRef<"form">) {
  return (
    <form className={cn("flex flex-col gap-6", className)} {...props}>
      <div className="flex flex-col items-start gap-2 text-left">
        <h1 className="text-4xl font-bold text-main font-helveticaBold">
          Register
        </h1>
        <p className="text-balance text-sm text-gray font-semibold">
          Register as new affiliate
        </p>
      </div>
      <div className="grid gap-6">
        <div className="grid gap-2">
          <Label htmlFor="email">Email</Label>
          <Input id="email" type="email" placeholder="m@example.com" required />
        </div>
        <div className="grid gap-2">
          <div className="flex items-center">
            <Label htmlFor="password">Password</Label>
          </div>
          <Input id="password" type="password" required />
        </div>

        <div className="grid gap-2">
          <div className="flex items-center">
            <Label htmlFor="confirmpassword">Confirm Password</Label>
          </div>
          <Input id="confirmpassword" type="password" required />
        </div>

        <Button type="submit" className="w-full h-12">
          Sign In
        </Button>
      </div>
      <div className="text-center text-sm font-semibold text-gray">
        Already have an account?{" "}
        <Link
          href="/auth/signin"
          className="underline underline-offset-4 text-main"
        >
          Sign in
        </Link>
      </div>
    </form>
  );
}
