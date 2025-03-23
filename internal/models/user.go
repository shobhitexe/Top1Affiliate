package models

type User struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	AffiliateID string `json:"affiliateId"`
	Password    string `json:"-"`
	Country     string `json:"country"`
	Commission  int    `json:"commission"`
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LeadsEmails struct {
	AffiliateID string
	Email       string
}
