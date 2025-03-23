package models

type Admin struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
}
