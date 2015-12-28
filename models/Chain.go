package models

// IPTablesSummary ..
type Chain struct {
	ID         int    `json:"id"`
	Policy     string `json:"policy"`
	Packets    string `json:"packets"`
	Bytes      string `json:"bytes"`
	References string `json:"references"`
	Rules      []Rule `json:"rules"`
}
