package domain_test

import (
	"strings"
	"testing"

	domain "github.com/kory-jp/vue_go/api/domain/account"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestTranslateUsersField(t *testing.T) {
	account := &domain.Account{}

	cases := []struct {
		name   string
		args   string
		result string
	}{
		{
			name:   "success",
			args:   "Name",
			result: "名前",
		},
		{
			name:   "success",
			args:   "Email",
			result: "メールアドレス",
		},
		{
			name:   "success",
			args:   "Password",
			result: "パスワード",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			result := account.TranslateUsersField(tt.args)
			assert.Equal(t, tt.result, result)
		})
	}
}

func TestAccountValidate(t *testing.T) {

	cases := []struct {
		name    string
		account *domain.Account
		err     error
	}{
		{
			name: "validate = success",
			account: &domain.Account{
				Name:     "testName",
				Email:    "test@exm.com",
				Password: "testPassword",
			},
			err: nil,
		},
		{
			name: "when name is empty = fail",
			account: &domain.Account{
				Email:    "test@exm.com",
				Password: "testPassword",
			},
			err: errors.New("名前は必須です。"),
		},
		{
			name: "when email is empty = fail",
			account: &domain.Account{
				Name:     "testName",
				Password: "testPassword",
			},
			err: errors.New("メールアドレスは必須です。"),
		},
		{
			name: "when password is empty = fail",
			account: &domain.Account{
				Name:  "testName",
				Email: "test@exm.com",
			},
			err: errors.New("パスワードは必須です。"),
		},
		{
			name: "when name<2 = fail",
			account: &domain.Account{
				Name:     "1",
				Email:    "test@exm.com",
				Password: "testPassword",
			},
			err: errors.New("名前は2文字より入力が必須です。"),
		},
		{
			name: "when name>19 = fail",
			account: &domain.Account{
				Name:     strings.Repeat("a", 20),
				Email:    "test@exm.com",
				Password: "testPassword",
			},
			err: errors.New("名前は20文字未満の入力になります。"),
		},
		{
			name: "when email>29 = fail",
			account: &domain.Account{
				Name:     "testName",
				Email:    strings.Repeat("a", 30) + "@exm.com",
				Password: "testPassword",
			},
			err: errors.New("メールアドレスは30文字未満の入力になります。"),
		},
		{
			name: "when Email is not formatted correctly= fail",
			account: &domain.Account{
				Name:     "testName",
				Email:    "aaaaaaaaa",
				Password: "testPassword",
			},
			err: errors.New("メールアドレスのフォーマットに誤りがあります"),
		},
		{
			name: "when password<5 = fail",
			account: &domain.Account{
				Name:     "testName",
				Email:    "test@exm.com",
				Password: "test",
			},
			err: errors.New("パスワードは5文字より入力が必須です。"),
		},
		{
			name: "when password>19 = fail",
			account: &domain.Account{
				Name:     "testName",
				Email:    "test@exm.com",
				Password: strings.Repeat("a", 20),
			},
			err: errors.New("パスワードは20文字未満の入力になります。"),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.account.AccountValidate()
			if err != nil {
				assert.Equal(t, tt.err.Error(), err.Error())
			}
		})
	}
}

func TestCheckExistEmail(t *testing.T) {
	account := &domain.Account{}
	cases := []struct {
		name    string
		args    int
		errArgs error
		err     error
	}{
		{
			name:    "success",
			args:    0,
			errArgs: nil,
			err:     nil,
		},
		{
			name:    "when email is exist = fail",
			args:    1,
			errArgs: nil,
			err:     errors.New("入力されたメールアドレスは既に登録されております"),
		},
		{
			name:    "when error is exist = fail",
			args:    0,
			errArgs: errors.New("error is exist"),
			err:     errors.New("error is exist"),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			err := account.CheckExistEmail(tt.args, tt.errArgs)
			if err != nil {
				assert.Equal(t, tt.err.Error(), err.Error())
			}
		})
	}
}
