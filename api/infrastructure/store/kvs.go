package store

import (
	"context"
	"fmt"
	"time"

	domain "github.com/kory-jp/vue_go/api/domain/account"

	"github.com/kory-jp/vue_go/api/config"
	"github.com/redis/go-redis/v9"
)

type KVS struct {
	Cli *redis.Client
}

func NewKVS() (*KVS, error) {
	cli := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", config.Config.RedisHost, config.Config.RedisPort),
	})
	return &KVS{Cli: cli}, nil
}

func (k *KVS) Save(ctx context.Context, key string, accountID int) error {
	return k.Cli.Set(ctx, key, accountID, 30*time.Minute).Err()
}

func (k *KVS) Load(ctx context.Context, key string) (account *domain.Account, err error) {
	id, err := k.Cli.Get(ctx, key).Int64()
	if err != nil {
		return nil, fmt.Errorf("failed to get by %q: %w", key, err)
	}
	account = &domain.Account{
		ID: int(id),
	}
	return account, nil
}

func (k *KVS) Expire(ctx context.Context, key string, minitue time.Duration) error {
	_, err := k.Cli.Expire(ctx, key, minitue*time.Minute).Result()
	if err != nil {
		return err
	}
	return nil
}
