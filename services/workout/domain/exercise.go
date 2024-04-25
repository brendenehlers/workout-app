package domain

type Exercise struct {
	Name          string `json:"name"`
	Duration      int    `json:"duration"`
	DurationUnits string `json:"durationUnits"`
	Sets          int    `json:"sets"`
}

type ExerciseData struct {
	Exercise
	Id        string    `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	Embedding Embedding `json:"-"`
}

type ExerciseRepository interface {
	CreateExercise(*Exercise) (*ExerciseData, error)
	GetExercise(string) (*ExerciseData, error)
	GetExercises() ([]*ExerciseData, error)
	UpdateExercise(string, *ExerciseData) (*ExerciseData, error)
	DeleteExercise(string) error
}
