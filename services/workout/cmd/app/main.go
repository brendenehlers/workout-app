package main

import (
	"github.com/brendenehlers/workout-app/services/workout/domain"
	"github.com/brendenehlers/workout-app/services/workout/http"
	"github.com/brendenehlers/workout-app/services/workout/log"
	"github.com/brendenehlers/workout-app/services/workout/search"
)

func main() {
	s := search.New(Embedder{}, VectorStore{})
	server := http.New(":8080", s)
	log.Fatal(server.Start())
}

type Embedder struct{}

func (Embedder) GetEmbedding(query string) (*domain.Embedding, error) {
	emb := domain.Embedding(make([]float32, 10))
	return &emb, nil
}

type VectorStore struct{}

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
