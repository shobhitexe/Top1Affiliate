package models

type Admin struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
}

type AddAffiliate struct {
	AffiliateID string `json:"affiliateid"`
	Name        string `json:"name"`
	Country     string `json:"country"`
	Commission  int    `json:"commission"`
	Password    string `json:"password"`
}

type EditAffiliate struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Country    string `json:"country"`
	Commission int    `json:"commission"`
}
