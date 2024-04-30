package domain

type VSDocument struct {
	Embedding *Embedding
	Data      *Exercise
}

type VectorStore interface {
	Nearby(*Embedding) ([]*VSDocument, error)
}
