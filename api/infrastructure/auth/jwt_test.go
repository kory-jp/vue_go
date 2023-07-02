package auth

import (
	"bytes"
	"context"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kory-jp/vue_go/api/utiles/fixture"

	"github.com/golang/mock/gomock"
	domain "github.com/kory-jp/vue_go/api/domain/account"
	mock_auth "github.com/kory-jp/vue_go/api/infrastructure/auth/mock"
)

func TestEmbed(t *testing.T) {
	want := []byte("-----BEGIN PUBLIC KEY-----")
	if !bytes.Contains(rawPubKey, want) {
		t.Errorf("want %s, but got %s", want, rawPubKey)
	}
	want = []byte("-----BEGIN RSA PRIVATE KEY-----")
	if !bytes.Contains(rawPrivKey, want) {
		t.Errorf("want %s, but got %s", want, rawPubKey)
	}
}

func TestJWTer_GenerateToken(t *testing.T) {
	ctx := new(gin.Context)
	c := gomock.NewController(t)
	defer c.Finish()
	ms := mock_auth.NewMockStore(c)

	cases := []struct {
		name               string
		args               *domain.Account
		prepareStoreMockFn func(*mock_auth.MockStore, context.Context, string, int)
	}{
		{
			name: "success",
			args: fixture.Account(
				&domain.Account{
					ID: 20,
				},
			),
			prepareStoreMockFn: func(m *mock_auth.MockStore, ctx context.Context, key string, accountID int) {
				m.EXPECT().Save(ctx, gomock.Any(), accountID).Return(nil)
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			sut, err := NewJWTer(ms)
			if err != nil {
				t.Fatal(err)
			}
			tt.prepareStoreMockFn(ms, ctx, string(rawPubKey), tt.args.ID)
			got, err := sut.GenerateToken(ctx, tt.args)
			if err != nil {
				t.Fatalf("not want err: %v", err)
			}
			if len(got) == 0 {
				t.Errorf("token is empty")
			}
		})
	}
}
