package usecase_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	domain "github.com/kory-jp/vue_go/api/domain/account"
	usecase "github.com/kory-jp/vue_go/api/usecase/account"
	mock_usecase "github.com/kory-jp/vue_go/api/usecase/account/mock"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()
	AccountRepository := mock_usecase.NewMockAccountRepository(c)
	inter := &usecase.AccountInteractor{}
	inter.AccountRepository = AccountRepository

	cases := []struct {
		name                     string
		args                     domain.Account
		prepareStoreMockFn       func(*mock_usecase.MockAccountRepository, domain.Account)
		prepareFindByIdMockFn    func(*mock_usecase.MockAccountRepository, int, *domain.Account)
		prepareFindByEmailMockFn func(*mock_usecase.MockAccountRepository, string)
		result                   *domain.Account
		err                      error
	}{
		{
			name: "add = success",
			args: domain.Account{
				Name:     "testName",
				Email:    "test@exm.com",
				Password: "testPassword",
			},
			prepareStoreMockFn: func(m *mock_usecase.MockAccountRepository, args domain.Account) {
				m.EXPECT().Store(gomock.Any()).Return(1, nil)
			},
			prepareFindByIdMockFn: func(m *mock_usecase.MockAccountRepository, id int, result *domain.Account) {
				m.EXPECT().FindById(id).Return(result, nil)
			},
			prepareFindByEmailMockFn: func(m *mock_usecase.MockAccountRepository, email string) {
				m.EXPECT().FindByEmail(email).Return(0, nil)
			},
			result: &domain.Account{
				ID:       1,
				Name:     "testUser",
				Email:    "test@exm.com",
				Password: "test_password",
			},
			err: nil,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepareStoreMockFn(AccountRepository, tt.args)
			tt.prepareFindByEmailMockFn(AccountRepository, tt.result.Email)
			tt.prepareFindByIdMockFn(AccountRepository, tt.result.ID, tt.result)
			_, err := inter.Add(tt.args)
			assert.Equal(t, tt.err, err)
		})
	}
}
