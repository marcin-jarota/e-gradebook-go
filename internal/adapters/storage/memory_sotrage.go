package storage

type MemoryStorage struct {
	data map[string]any
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: map[string]any{},
	}
}

func (s *MemoryStorage) Get(key string) (any, error) {
	return s.data[key], nil
}

func (s *MemoryStorage) Set(key string, value any) error {
	s.data[key] = value
	return nil
}

func (s *MemoryStorage) Delete(key string) error {
	delete(s.data, key)
	return nil
}
