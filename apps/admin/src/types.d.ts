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

type Affiliate = {
  id: string;
  name: string;
  affiliateId: string;
  country: string;
  commission: number;
  link: string;
};

type WeeklyStatsData = {
  registrations: number;
  deposits: number;
  withdrawals: number;
  commission: number;
};

type StatsData = {
  registrations: number;
  deposits: number;
  withdrawals: number;
  commission: number;

  registrationsMonthly: number;
  depositsMonthly: number;
  withdrawalsMonthly: number;
  commissionMonthly: number;
};

type AffiliatePathType = {
  id: string;
  name: string;
  addedBy: string;
};
