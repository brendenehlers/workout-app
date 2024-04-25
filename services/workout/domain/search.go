package domain

type Embedding []float32

type SearchService interface {
	Search(WorkoutQuery) (*WorkoutData, error)
}
