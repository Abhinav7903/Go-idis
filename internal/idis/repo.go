package idis

type Repository interface {
	Set(key string, value string) error
	Get(key string) (string, error)
	Delete(key string) error
}