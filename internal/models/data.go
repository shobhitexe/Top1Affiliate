package models

type Leads struct {
	ID                   int    `json:"id"`
	FirstName            string `json:"firstName"`
	LastName             string `json:"lastName"`
	Updated              string `json:"updated"`
	LastLoginDate        string `json:"lastLoginDate"`
	LeadGuid             string `json:"leadGuid"`
	Country              string `json:"country"`
	City                 string `json:"city"`
	TimeZone             string `json:"timeZone"`
	SalesStatus          string `json:"salesStatus"`
	Language             string `json:"language"`
	BusinessUnit         string `json:"businessUnit"`
	DomainName           string `json:"domainName"`
	IsQualified          bool   `json:"isQualified"`
	ConversionAgentID    int    `json:"conversionAgentId"`
	RetentionManagerID   int    `json:"retentionManagerId"`
	VIPManagerID         int    `json:"vipManagerId"`
	CloserManagerID      int    `json:"closerManagerId"`
	ConversionAgentTeam  string `json:"conversionAgentTeam"`
	RetentionManagerTeam string `json:"retentionManagerTeam"`
	VIPManagerTeam       string `json:"vipManagerTeam"`
	CloserManagerTeam    string `json:"closerManagerTeam"`
	AffiliateID          string `json:"affiliateId"`
	AffiliateName        string `json:"affiliateName"`
	UTMCampaign          string `json:"utmCampaign"`
	UTMMedium            string `json:"utmMedium"`
	UTMSource            string `json:"utmSource"`
	UTMTerm              string `json:"utmTerm"`
	ReferringPage        string `json:"referringPage"`
	RegistrationDate     string `json:"registrationDate"`
	AccountCreationDate  string `json:"accountCreationDate"`
	ActivationDate       string `json:"activationDate"`
	FullyActivationDate  string `json:"fullyActivationDate"`
	SubChannel           string `json:"subChannel"`
	ChannelName          string `json:"channelName"`
	TLName               string `json:"tlName"`
	TrackingLinkID       string `json:"trackingLinkId"`
	Deposited            bool   `json:"deposited"`
	OriginalLeadID       int    `json:"originalLeadId"`
	OriginalByNameLeadID int    `json:"originalByNameLeadId"`
	NameDuplicates       string `json:"nameDuplicates"`
	Email                string `json:"email"`
	OfferDescription     string `json:"offerDescription"`
	IPAddress            string `json:"ipAddress"`
	LandingPage          string `json:"landingPage"`
}

type Transaction struct {
	TransactionID      int     `json:"transactionId"`
	Amount             float64 `json:"amount"`
	TransactionType    string  `json:"transactionType"`
	TransactionSubType string  `json:"transactionSubType"`
	Status             string  `json:"status"`
	TransactionDate    string  `json:"transactionDate"`
	LeadID             int     `json:"leadId"`
	LeadGUID           string  `json:"leadGuid"`
}

type Stats struct {
	Registrations int     `json:"registrations"`
	Deposits      float64 `json:"deposits"`
	Withdrawals   float64 `json:"withdrawals"`
	Commissions   float64 `json:"commission"`
}

type WeeklyStatsWithMonthly struct {
	Registrations int     `json:"registrations"`
	Deposits      float64 `json:"deposits"`
	Withdrawals   float64 `json:"withdrawals"`
	Commissions   float64 `json:"commission"`

	RegistrationsMonthly int     `json:"registrationsMonthly"`
	DepositsMonthly      float64 `json:"depositsMonthly"`
	WithdrawalsMonthly   float64 `json:"withdrawalsMonthly"`
	CommissionsMonthly   float64 `json:"commissionMonthly"`
}

type CommissionTxn struct {
	LeadID  string  `json:"id"`
	Name    string  `json:"name"`
	Country string  `json:"country"`
	Email   string  `json:"email"`
	Date    string  `json:"date"`
	Amount  float64 `json:"amount"`
	TxnType string  `json:"txnType"`
}

type DashboardStats struct {
	Weekly      Stats           `json:"weekly"`
	Commissions []CommissionTxn `json:"commissions"`
}

type Leaderboard struct {
	AffiliateId      string  `json:"affiliateId"`
	Name             string  `json:"name"`
	Country          string  `json:"country"`
	TotalCommissions float64 `json:"totalCommissions"`
}

type Statistics struct {
	AffiliateID      string  `json:"affiliateId"`
	FirstName        string  `json:"firstName"`
	LastName         string  `json:"lastName"`
	Country          string  `json:"country"`
	RegistrationDate string  `json:"registrationDate"`
	Deposits         float64 `json:"deposits"`
	Withdrawals      float64 `json:"withdrawals"`
	Commissions      float64 `json:"commission"`
}
