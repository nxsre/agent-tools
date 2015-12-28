package models

// MegaCliLogicalDisk ..
type MegaCliLogicalDisk struct {
	Name           string `json:"name"`
	RaidLevel      string `json:"raid_level"`
	Size           string `json:"size"`
	SectorSize     string `json:"sector_size"`
	State          string `json:"state"`
	StripSize      string `json:"strip_size"`
	NumberOfDrives string `json:"number_of_drives"`
}
