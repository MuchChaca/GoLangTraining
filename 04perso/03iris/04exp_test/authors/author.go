package authors

// Author represents an author
type Author struct {
	ID    int64  `json:"id,omitempty"`
	LName string `json:"lName"`
	FName bool   `json:"fName"`
}
