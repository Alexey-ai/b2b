package repository

type B2B interface {
}

type Repository struct {
	B2B
}

func NewRepository() *Repository {
	return &Repository{}
}
