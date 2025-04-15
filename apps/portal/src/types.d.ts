// eslint-disable-next-line @typescript-eslint/no-unused-vars
import NextAuth from "next-auth";

type sessionUser = {
  id?: string;
  affiliateId?: string;
  commission?: number;
  Clientlink?: string;
  Sublink?: string;

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

type Leads = {
  id: number;
  firstName: string;
  lastName: string;
  updated: string;
  lastLoginDate: string;
  leadGuid: string;
  country: string;
  city: string;
  timeZone: string;
  salesStatus: string;
  language: string;
  businessUnit: string;
  domainName: string;
  isQualified: boolean;
  conversionAgentId: number;
  retentionManagerId: number;
  vipManagerId: number;
  closerManagerId: number;
  conversionAgentTeam: string;
  retentionManagerTeam: string;
  vipManagerTeam: string;
  closerManagerTeam: string;
  affiliateId: string;
  affiliateName: string;
  utmCampaign: string;
  utmMedium: string;
  utmSource: string;
  utmTerm: string;
  referringPage: string;
  registrationDate: string;
  accountCreationDate: string;
  activationDate: string;
  fullyActivationDate: string;
  subChannel: string;
  channelName: string;
  tlName: string;
  trackingLinkId: string;
  deposited: boolean;
  originalLeadId: number;
  originalByNameLeadId: number;
  nameDuplicates: string;
  email: string;
  offerDescription: string;
  ipAddress: string;
  landingPage: string;
};

type CommissionTxn = {
  id: string;
  name: string;
  country: string;
  email: string;
  date: string;
  amount: number;
  txnType: string;
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

type WeeklyStatsData = {
  registrations: number;
  deposits: number;
  withdrawals: number;
  commission: number;
  ftds: number;
};

type DashboardStats = {
  weekly: WeeklyStatsData;
  net: WeeklyStatsData;
  commissions: CommissionTxn[];
  sales: MonthlySalesOverview[];
  commissionStats: MonthlyCommissionOverview[];
};

type WalletDetails = {
  iban: string;
  swift: string;
  bankName: string;
  chainName: string;
  walletAddress: string;
};

type AffiliatePathType = {
  id: string;
  name: string;
  addedBy: string;
};

type MonthlySalesOverview = {
  month: string;
  deposits: number;
  withdrawals: number;
};

type MonthlyCommissionOverview = {
  month: string;
  commission: number;
};
