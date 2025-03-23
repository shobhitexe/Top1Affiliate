import { BackendURL } from "@/config/env";
import type { NextAuthOptions } from "next-auth";
import CredentialsProvider from "next-auth/providers/credentials";

export const options: NextAuthOptions = {
  providers: [
    CredentialsProvider({
      name: "Credentials",
      credentials: {
        login: { label: "Login", type: "text" },
        password: { label: "Password", type: "password" },
      },
      async authorize(credentials) {
        const response = await fetch(`${BackendURL}/api/v1/admin/login`, {
          method: "POST",
          body: JSON.stringify({ ...credentials }),
          headers: { "Content-Type": "application/json" },
          credentials: "include",
        });

        if (response.status === 200) {
          const data = await response.json();

          return data.data;
        } else {
          return null;
        }
      },
    }),
  ],
  session: { strategy: "jwt" },
  callbacks: {
    async jwt({ token, user, trigger, session }) {
      if (trigger === "update") {
        return { ...token, ...session.user };
      }
      return { ...token, ...user };
    },
    async session({ session, token }) {
      session.user = token;

      return session;
    },
  },
  pages: {
    signIn: "/auth/signin",
  },
};
