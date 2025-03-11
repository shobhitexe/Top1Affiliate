package models

type User struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	AffiliateID string `json:"affiliateId"`
	Password    string `json:"-"`
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
