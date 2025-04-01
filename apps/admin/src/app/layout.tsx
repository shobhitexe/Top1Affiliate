import type { Metadata } from "next";
import "./globals.css";
import { getServerSession } from "next-auth";
import { options } from "./api/auth/[...nextauth]/options";
import { SessionProviders } from "@/components";
import { Toaster } from "@/components/ui/toaster";

export const metadata: Metadata = {
  title: "Top1 | Admin",
  description: "",
};

export default async function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const session = await getServerSession(options);

  return (
    <html lang="en">
      <SessionProviders session={session}>
        <body className={`antialiased`}>
          {children} <Toaster />
        </body>
      </SessionProviders>
    </html>
  );
}
