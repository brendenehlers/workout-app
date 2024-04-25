package domain

type CreateWorkoutRequest struct {
	Query    string   `json:"query"`
	Excludes []string `json:"excludes"`
}
