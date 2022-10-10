package repository

type Test interface {
}

type Repository struct {
	Test
}

func NewRepository() *Repository {
	return &Repository{}
}
