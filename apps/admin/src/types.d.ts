// eslint-disable-next-line @typescript-eslint/no-unused-vars
import NextAuth from "next-auth";

type sessionUser = {
  id?: string;
  username?: string;

  name?: string | null | undefined;
  email?: string | null | undefined;
};

declare module "next-auth" {
  interface Session {
    user: sessionUser;
  }
  interface nextauth {
    token: string;
  }
}
