package models

// IPTablesSummary ..
type Table struct {
	Chains map[string]Chain `json:"chains"`
}
