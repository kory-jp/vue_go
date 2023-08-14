package auth

import (
	"context"
	_ "embed"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/kory-jp/vue_go/api/infrastructure/clock"
	"github.com/kory-jp/vue_go/api/interfaces/controllers"

	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwa"

	"github.com/gin-gonic/gin"

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
	Load(ctx context.Context, key string) (*domain.Account, error)
	Expire(ctx context.Context, key string, minitue time.Duration) error
	Delete(ctx context.Context, key string) error
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
	AccountID = "account_id"
	Email     = "email"
)

func (j *JWTer) GenerateToken(ctx controllers.Context, account *domain.Account) ([]byte, error) {
	tok, err := jwt.NewBuilder().
		JwtID(uuid.New().String()).
		Issuer(`github.com/kory-jp/vue_go/api`).
		Subject("access_token").
		IssuedAt(j.Clocker.Now()).
		Expiration(j.Clocker.Now().Add(30*time.Minute)).
		Claim(Email, account.Email).
		Claim(AccountID, account.ID).
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

func (j *JWTer) GetToken(ctx controllers.Context, r *http.Request) (jwt.Token, error) {
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

type accountID struct{}

func (j *JWTer) SetAccountID(ctx context.Context, aid int) context.Context {
	return context.WithValue(ctx, accountID{}, aid)
}

func (j *JWTer) GetAccountID(ctx controllers.Context) int {
	aid := ctx.Value(AccountID)
	v, ok := aid.(int64)
	if !ok {
		return 0
	}
	return int(v)
}

func (j *JWTer) FillContxet(ctx *gin.Context) error {
	token, err := j.GetToken(ctx, ctx.Request)
	if err != nil {
		return err
	}
	if err := jwt.Validate(token, jwt.WithClock(j.Clocker)); err != nil {
		return fmt.Errorf("GetToken: failed to validate token: %w", err)
	}
	id, ok := token.Get(AccountID)
	if !ok {
		return fmt.Errorf("not found %s", AccountID)
	}
	aid := fmt.Sprintf("%v", id)
	jwi, err := j.Store.Load(ctx, token.JwtID())
	if err != nil {
		return err
	}
	jwiID := fmt.Sprintf("%v", jwi.ID)
	if jwiID != aid {
		return fmt.Errorf("expired token %s because login another", string(rune(jwi.ID)))
	}
	err = j.Store.Expire(ctx, aid, time.Duration(60))
	if err != nil {
		return err
	}
	intAid, err := strconv.ParseInt(aid, 10, 64)
	if err != nil {
		return nil
	}
	ctx.Set(AccountID, intAid)

	get, ok := ctx.Get(Email)
	if !ok {
		ctx.Set(Email, "")
		return nil
	}
	ctx.Set(Email, get)

	return nil
}

func (j *JWTer) DeleteAccountID(ctx controllers.Context, key string) error {
	if err := j.Store.Delete(ctx, key); err != nil {
		return err
	}
	return nil
}
