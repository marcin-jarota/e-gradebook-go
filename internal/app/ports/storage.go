package ports

type SessionStorage interface {
	Get(key string) (any, error)
	Set(key string, value any) error
	Delete(key string) error
}
