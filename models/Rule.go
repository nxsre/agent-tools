package models

// Rule ..
type Rule struct {
	Num         string `json:"num"`
	Pkts        string `json:"pkts"`
	Bytes       string `json:"bytes"`
	Target      string `json:"target"`
	Prot        string `json:"prot"`
	Opt         string `json:"opt"`
	In          string `json:"in"`
	Out         string `json:"out"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Details     string `json:"details"`
	Comment     string `json:"comment"`
}
