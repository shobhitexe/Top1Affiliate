package models

type User struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	AffiliateID string  `json:"affiliateId"`
	Balance     float64 `json:"balance"`
	Password    string  `json:"-"`
	Country     string  `json:"country"`
	Commission  int     `json:"commission"`
	Blocked     bool    `json:"blocked"`
	ClientLink  string  `json:"Clientlink"`
	SubLink     string  `json:"Sublink"`
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LeadsEmails struct {
	AffiliateID string
	Email       string
}

type Payouts struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	AffiliateId string  `json:"affiliateId"`
	Amount      float64 `json:"amount"`
	Type        string  `json:"type"`
	Status      string  `json:"status"`
	Method      string  `json:"method"`
	CreatedAt   string  `json:"createdAt"`

	IBAN          string `json:"iban"`
	SwiftCode     string `json:"swiftCode"`
	BankName      string `json:"bankName"`
	ChainName     string `json:"chainName"`
	WalletAddress string `json:"walletAddress"`
}

type RequestPayout struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
	Type   string  `json:"type"`
	Method string  `json:"method"`
}

type WalletDetails struct {
	ID            string `json:"id"`
	IBAN          string `json:"iban"`
	Swift         string `json:"swift"`
	BankName      string `json:"bankName"`
	ChainName     string `json:"chainName"`
	WalletAddress string `json:"walletAddress"`
}
