package domain

type Exercise struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}
