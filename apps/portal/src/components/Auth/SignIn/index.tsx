import { cn } from "@/lib/utils";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import Link from "next/link";
import { Switch } from "@/components/ui/switch";

export function LoginForm({
  className,
  ...props
}: React.ComponentPropsWithoutRef<"form">) {
  return (
    <form className={cn("flex flex-col gap-6", className)} {...props}>
      <div className="flex flex-col items-start gap-2 text-left">
        <h1 className="text-4xl font-bold text-main font-helveticaBold">
          Welcome Back
        </h1>
        <p className="text-balance text-sm text-gray font-semibold">
          Enter your email and password to sign in
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
            {/* <a
              href="#"
              className="ml-auto text-sm underline-offset-4 hover:underline"
            >
              Forgot your password?
            </a> */}
          </div>
          <Input id="password" type="password" required />
        </div>

        <div className="flex items-center gap-3">
          <Switch />
          <Label htmlFor="remember">Remember Me</Label>
        </div>

        <Button type="submit" className="w-full h-12">
          Sign In
        </Button>
      </div>
      <div className="text-center text-sm font-semibold text-gray">
        Don&apos;t have an account?{" "}
        <Link
          href="/auth/signup"
          className="underline underline-offset-4 text-main"
        >
          Sign up
        </Link>
      </div>
    </form>
  );
}
