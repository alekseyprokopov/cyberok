package model

type Whois struct {
	domain string `json:"domain" db:"domain"`
	whois  string `json:"getWhois" db:"getWhois"`
}
