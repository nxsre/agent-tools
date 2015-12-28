package models

// BGPNextHop ..
type BGPNextHop struct {
	IP   string `json:"ip"`
	AFI  string `json:"afi"`
	Used bool   `json:"used"`
}
