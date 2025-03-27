package models

type User struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	AffiliateID string `json:"affiliateId"`
	Password    string `json:"-"`
	Country     string `json:"country"`
	Commission  int    `json:"commission"`
	Blocked     bool   `json:"blocked"`
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
	Name        string  `json:"name"`
	AffiliateId string  `json:"affiliateId"`
	Amount      float64 `json:"amount"`
	Type        string  `json:"type"`
	Status      string  `json:"status"`
	CreatedAt   string  `json:"createdAt"`
}

type RequestPayout struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
	Type   string  `json:"type"`
}
