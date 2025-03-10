// eslint-disable-next-line @typescript-eslint/no-unused-vars
import NextAuth from "next-auth";

type sessionUser = {
  id?: string;
  cookie?: string;

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
