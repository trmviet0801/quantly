package repos

type BaseRepository[T any] interface {
	GetById(id string) (*T, error)
	Create(data *T) error
	Update(data *T) error
	DeleteById(id string) error
}
