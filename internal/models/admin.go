package models

type Admin struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
}

type AddAffiliate struct {
	AddedBy     int    `json:"addedBy"`
	AffiliateID string `json:"affiliateid"`
	Name        string `json:"name"`
	Country     string `json:"country"`
	Commission  int    `json:"commission"`
	Password    string `json:"password"`
	ClientLink  string `json:"Clientlink"`
	SubLink     string `json:"Sublink"`
}

type EditAffiliate struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Country    string  `json:"country"`
	Commission int     `json:"commission"`
	ClientLink string  `json:"Clientlink"`
	SubLink    string  `json:"Sublink"`
	Balance    float64 `json:"balance"`
}

type AffiliatePath struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	AddedBy string `json:"addedBy"`
	Depth   string `json:"depth"`
}
