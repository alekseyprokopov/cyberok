package model

type Fqdn struct {
	Name string   `json:"fqdn" db:"fqdn_name"`
	Ips  []string `json:"ips" db:"ip"`
}

// TODO
func (s Fqdn) isValid() bool {
	return true
}
