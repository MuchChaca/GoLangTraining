package genres

// Genre represents a genre
type Genre struct {
	// SessionID string `json:"-"`
	ID    int64  `json:"id,omitempty"`
	Label string `json:"label"`
}
