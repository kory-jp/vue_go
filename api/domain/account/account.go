package domain

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" validate:"required,gte=2,lt=20"`
	Email     string    `json:"email" validate:"required,email,lt=30"`
	Password  string    `json:"password" validate:"required,gte=5,lt=20"`
	CreatedAt time.Time `json:"createdAt"`
}

func (ac Account) Encrypt(plaintext string) (hash string) {
	byteHash, _ := bcrypt.GenerateFromPassword([]byte(plaintext), bcrypt.DefaultCost)
	hash = string(byteHash)
	return hash
}

func (ac Account) TranslateUsersField(field string) (value string) {
	switch field {
	case "Name":
		value = "名前"
	case "Email":
		value = "メールアドレス"
	case "Password":
		value = "パスワード"
	}
	return
}

func (ac *Account) AccountValidate() (err error) {
	validate := validator.New()
	err = validate.Struct(ac)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			value := ac.TranslateUsersField(err.Field())
			switch err.ActualTag() {
			case "required":
				return errors.New(fmt.Sprintf("%sは必須です。", value))
			case "gte":
				return errors.New(fmt.Sprintf("%sは%s文字より入力が必須です。", value, err.Param()))
			case "lt":
				return errors.New(fmt.Sprintf("%sは%s文字未満の入力になります。", value, err.Param()))
			case "email":
				return errors.New("メールアドレスのフォーマットに誤りがあります")
			}
		}
	}
	return err
}

func (ac *Account) CheckExistEmail(numberAccount int, err error) error {
	if err != nil {
		return err
	}
	if numberAccount > 0 {
		return errors.New("入力されたメールアドレスは既に登録されております")
	}
	return nil
}
