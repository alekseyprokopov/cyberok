package model

type Whois struct {
	Domain string `json:"domain" db:"domain"`
	Whois  string `json:"info" db:"info"`
}
