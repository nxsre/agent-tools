package models

// IPTablesSummary ..
type IPTablesSummary struct {
	Tables map[string]Table `json:"tables"`
}
