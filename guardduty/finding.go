package guardduty

// Finding represents a single gd finding
type Finding struct {
	ThreatPurpose string
	Data          Metadata
}

type Metadata struct {
	count int `json:"count"`

	// countless other fields...
}

// NewFinding is the finding constructor
func NewFinding(tp string) *Finding {
	return &Finding{
		ThreatPurpose: tp,
		Data:          Metadata{count: 1},
	}
}

// Inc increments the count on a finding
func (f *Finding) Inc() {
	f.Data.count++
}
