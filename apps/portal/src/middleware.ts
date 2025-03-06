import { withAuth } from "next-auth/middleware";

export default withAuth(
  async function middleware() {},

  {
    callbacks: {
      authorized: ({ req, token }) =>
        req.nextUrl.pathname.startsWith("/auth/") || !!token,
    },

    pages: {
      signIn: "/auth/signin",
    },
  }
);

export const config = {
  matcher: [
    "/((?!api|_next/static|_next/image|images|favicon|fonts|favicon.ico).*)",
  ],
};
