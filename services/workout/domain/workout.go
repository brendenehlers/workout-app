package domain

type Workout struct {
	Name        string `json:"name"`
	Description string `json:"description"`

	Exercises []*Exercise `json:"exercises"`
}

type WorkoutData struct {
	*Workout
	Id string `json:"id"`
}

type WorkoutQuery struct {
	Query string `json:"query"`
}
