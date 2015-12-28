package models

// BGPRoute ..
type BGPRoute struct {
	Valid     bool         `json:"valid"`
	BestPath  bool         `json:"bestpath"`
	PathFrom  string       `json:"pathFrom"`
	MED       int          `json:"med"`
	Weight    int          `json:"weight"`
	LocalPref int          `json:"localpref"`
	ASPath    string       `json:"aspath"`
	Origin    string       `json:"origin"`
	NextHops  []BGPNextHop `json:"nexthops"`
	//Multipath bool       `json:"multipath"`
}
