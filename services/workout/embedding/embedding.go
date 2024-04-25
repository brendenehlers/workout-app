package embedding

import "github.com/brendenehlers/workout-app/services/workout/domain"

type Embedding struct{}

func New() *Embedding {
	return &Embedding{}
}

func (Embedding) GetEmbedding(query string) (*domain.Embedding, error) {
	emb := domain.Embedding(make([]float32, 10))
	return &emb, nil
}
