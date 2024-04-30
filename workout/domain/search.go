package domain

type SearchService interface {
	Search(WorkoutQuery) (*WorkoutData, error)
}
