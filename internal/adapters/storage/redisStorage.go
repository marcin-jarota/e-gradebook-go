package storage

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisStorage struct {
	client *redis.Client
	ctx    context.Context
	prefix string
}

func NewRedisStorage(prefix string, client *redis.Client, ctx context.Context) *RedisStorage {
	return &RedisStorage{
		client: client,
		ctx:    ctx,
		prefix: prefix,
	}
}

func (s *RedisStorage) Get(key string) (any, error) {
	val, err := s.client.Get(s.ctx, s.key(key)).Result()

	if err != nil {
		return nil, err
	}

	return val, nil
}

func (s *RedisStorage) Set(key string, value any) error {
	return s.client.Set(s.ctx, s.key(key), value, 0).Err()
}

func (s *RedisStorage) Delete(key string) error {
	return s.client.Del(s.ctx, s.key(key)).Err()
}

func (s *RedisStorage) key(baseKey string) string {
	return fmt.Sprintf("%s_%s", s.prefix, baseKey)
}
