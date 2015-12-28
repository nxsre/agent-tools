package models

// BGPRIB ..
type BGPRIB struct {
	TableVersion int                   `json:"tableVersion"`
	RouterId     string                `json:"routerId"`
	Routes       map[string][]BGPRoute `json:"routes"`
}
