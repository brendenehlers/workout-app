package vectorstore

import "github.com/brendenehlers/workout-app/services/workout/domain"

type VectorStore struct{}

func New() *VectorStore {
	return &VectorStore{}
}

func (VectorStore) Nearby(_ *domain.Embedding) ([]*domain.VSDocument, error) {
	return []*domain.VSDocument{
		{
			Embedding: &domain.Embedding{},
			Data: &domain.Exercise{
				Name:        "Lats",
				Description: "This works out your lats",
			},
		},
		{
			Embedding: &domain.Embedding{},
			Data: &domain.Exercise{
				Name: "Lats",
			},
		},
		{
			Embedding: &domain.Embedding{},
			Data: &domain.Exercise{
				Name: "Lats",
			},
		},
	}, nil
}
