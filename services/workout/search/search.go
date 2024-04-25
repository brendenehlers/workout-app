package search

import "github.com/brendenehlers/workout-app/services/workout/domain"

type SearchService struct {
	e  domain.EmbeddingService
	vs domain.VectorStore
}

func New(e domain.EmbeddingService, vs domain.VectorStore) *SearchService {
	return &SearchService{
		e:  e,
		vs: vs,
	}
}

func (s *SearchService) Search(q domain.WorkoutQuery) (*domain.WorkoutData, error) {
	emb, err := s.e.GetEmbedding(q.Query)
	if err != nil {
		return nil, err
	}

	docs, err := s.vs.Nearby(emb)
	if err != nil {
		return nil, err
	}

	workout := &domain.WorkoutData{
		Id: "new-workout",
		Workout: &domain.Workout{
			Name:        "Lat Blaster 5000",
			Description: "This will destroy your lats",
			Exercises:   make([]*domain.Exercise, 0, 10),
		},
	}

	for _, doc := range docs {
		workout.Exercises = append(workout.Exercises, doc.Data)
	}

	return workout, nil
}
