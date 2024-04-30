package domain

type Embedding []float32

type EmbeddingService interface {
	GetEmbedding(string) (*Embedding, error)
}
