package store_test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	domain "github.com/kory-jp/vue_go/api/domain/account"
	"github.com/kory-jp/vue_go/api/infrastructure/store"
	"github.com/redis/go-redis/v9"
)

func TestSave(t *testing.T) {
	t.Parallel()

	cli := OpenRedisForTest(t)

	sut := store.KVS{Cli: cli}
	key := "TestKVS_Save"
	account := domain.Account{ID: 1234}
	ctx := context.Background()
	t.Cleanup(func() {
		cli.Del(ctx, key)
	})
	if err := sut.Save(ctx, key, account.ID); err != nil {
		t.Errorf("want no error, but got %v", err)
	}
}

func TestLoad(t *testing.T) {
	t.Parallel()

	cli := OpenRedisForTest(t)
	sut := store.KVS{Cli: cli}

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		key := "TestKVS_Load_ok"
		account := domain.Account{ID: 1234}
		ctx := context.Background()
		cli.Set(ctx, key, int64(account.ID), 30*time.Minute)
		t.Cleanup(func() {
			cli.Del(ctx, key)
		})
		got, err := sut.Load(ctx, key)
		if err != nil {
			t.Fatalf("want no error, but got %v", err)
		}
		if got.ID != account.ID {
			t.Errorf("want %d, but got %d", account.ID, got.ID)
		}
	})

	t.Run("notFound", func(t *testing.T) {
		t.Parallel()

		key := "TestKVS_Save_notFound"
		ctx := context.Background()
		got, err := sut.Load(ctx, key)
		errMsg := "failed to get by \"TestKVS_Save_notFound\": redis: nil"
		if err != nil && err.Error() != errMsg {
			t.Errorf("want %v, but got %v(value = %v)", errors.New("TestKVS_Save_notFound"), err, got)
		}
	})
}

func OpenRedisForTest(t *testing.T) *redis.Client {
	t.Helper()

	host := "redis"
	port := 6379
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: "",
		DB:       0,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		t.Fatalf("failed to connect redis: %s", err)
	}
	return client
}
