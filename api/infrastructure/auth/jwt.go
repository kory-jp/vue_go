package auth

import (
	"context"
	_ "embed"
	"fmt"
	"net/http"
	"time"

	"github.com/kory-jp/vue_go/api/infrastructure/clock"

	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwa"

	domain "github.com/kory-jp/vue_go/api/domain/account"

	"github.com/pkg/errors"

	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

// 半角スペース禁止
//go:embed cert/secret.pem
var rawPrivKey []byte

//go:embed cert/public.pem
var rawPubKey []byte

type JWTer struct {
	PrivateKey, PublicKey jwk.Key
	Store                 Store
	Clocker               clock.Clocker
}

type Store interface {
	Save(ctx context.Context, key string, accountID int) error
	Load(ctx context.Context, key string) (accountID string, err error)
}

func NewJWTer(s Store) (*JWTer, error) {
	j := &JWTer{Store: s}
	privkey, err := parse(rawPrivKey)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	pubkey, err := parse(rawPubKey)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	j.PrivateKey = privkey
	j.PublicKey = pubkey
	j.Clocker = clock.RealClocker{}
	return j, nil
}

func parse(rawKey []byte) (jwk.Key, error) {
	key, err := jwk.ParseKey(rawKey, jwk.WithPEM(true))
	if err != nil {
		return nil, err
	}
	return key, nil
}

const (
	RoleKey        = "role"
	AccountNameKye = "account_name"
)

func (j *JWTer) GenerateToken(ctx context.Context, account domain.Account) ([]byte, error) {
	tok, err := jwt.NewBuilder().
		JwtID(uuid.New().String()).
		Issuer(`github.com/kory-jp/vue_go/api`).
		Subject("access_token").
		IssuedAt(j.Clocker.Now()).
		Expiration(j.Clocker.Now().Add(30*time.Minute)).
		Claim(AccountNameKye, account.Name).
		Build()
	if err != nil {
		return nil, fmt.Errorf("GetToken: failed to build token: %w", err)
	}
	if err := j.Store.Save(ctx, tok.JwtID(), account.ID); err != nil {
		return nil, err
	}
	signed, err := jwt.Sign(tok, jwt.WithKey(jwa.RS256, j.PrivateKey))
	if err != nil {
		return nil, err
	}
	return signed, nil
}

func (j *JWTer) GetToken(ctx context.Context, r *http.Request) (jwt.Token, error) {
	token, err := jwt.ParseRequest(
		r,
		jwt.WithKey(jwa.RS256, j.PublicKey),
		jwt.WithValidate(false),
	)
	if err != nil {
		return nil, err
	}
	if err := jwt.Validate(token, jwt.WithClock(j.Clocker)); err != nil {
		return nil, fmt.Errorf("get token: falied to validate token: %w", err)
	}

	if _, err := j.Store.Load(ctx, token.JwtID()); err != nil {
		return nil, fmt.Errorf("GetToken: %q expired: %w", token.JwtID(), err)
	}
	return token, nil
}
