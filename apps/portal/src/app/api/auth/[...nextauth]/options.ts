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
        const response = await fetch(`https://publicapi.fxlvls.com/login`, {
          method: "POST",
          body: JSON.stringify({ ...credentials }),
          headers: { "Content-Type": "application/json" },
          credentials: "include",
        });

        const resCookie = response.headers.getSetCookie();

        if (resCookie) {
          const cookieValue = resCookie[0].split(";")[0];

          return { cookie: cookieValue, id: credentials?.login ?? "" };
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
