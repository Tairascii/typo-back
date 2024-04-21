package repository

type Auth interface {
}

type Repository struct {
	Auth
}

func NewRepository() *Repository {
	return &Repository{}
}
