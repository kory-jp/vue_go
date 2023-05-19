package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	mock_controllers "github.com/kory-jp/vue_go/api/interfaces/controllers/account/mock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"github.com/kory-jp/vue_go/api/interfaces/database/account/mysql"

	controllers "github.com/kory-jp/vue_go/api/interfaces/controllers/account"
	mock_database "github.com/kory-jp/vue_go/api/interfaces/mock"

	"github.com/golang/mock/gomock"
	domain "github.com/kory-jp/vue_go/api/domain/account"
)

func TestMain(m *testing.M) {
	os.Setenv("PROCESS_ENV", "testProcess")
	os.Exit(m.Run())
}

func TestCreateDebug(t *testing.T) {
	var (
		err          error
		expectNumber int
		id           int
		name         string
		email        string
		password     string
		created_at   time.Time
	)
	c := gomock.NewController(t)
	defer c.Finish()
	sqlhandler := mock_database.NewMockSqlHandler(c)
	mock := mock_database.NewMockRow(c)
	mock2 := mock_database.NewMockRow(c)
	result := mock_database.NewMockResult(c)
	// ctx := context.Background()
	// kvs, err := store.NewKVS(ctx)
	// if err != nil {
	// 	log.Printf("[ERROR]: %+v", err)
	// }
	// fmt.Println(kvs)
	ctrl := controllers.NewAccountController(sqlhandler)
	var req *http.Request
	ac := domain.Account{
		Name:     "test_user",
		Email:    "test@exm.com",
		Password: "test_password",
	}
	jsonData, _ := json.Marshal(ac)
	req = httptest.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
	// --- Store ---
	result.EXPECT().LastInsertId().Return(int64(1), nil)
	sqlhandler.EXPECT().Execute(mysql.CreateAccountState, ac.Name, ac.Email, gomock.Any()).Return(result, nil)
	// --- FindByEmail ---
	mock.EXPECT().Next().Return(true)
	mock.EXPECT().Scan(&expectNumber).Return(nil)
	mock.EXPECT().Next().Return(false)
	mock.EXPECT().Err().Return(nil)
	defer mock.EXPECT().Close().Return(nil)
	sqlhandler.EXPECT().Query(mysql.GetNumberAccountState, ac.Email).Return(mock, err)
	// --- FindById ---
	defer mock2.EXPECT().Close().Return(nil)
	mock2.EXPECT().Next().Return(true)
	mock2.EXPECT().Scan(&id, &name, &email, &password, &created_at).Return(nil)
	mock2.EXPECT().Next().Return(false)
	mock2.EXPECT().Err().Return(nil)
	sqlhandler.EXPECT().Query(mysql.FindAccountState, 1).Return(mock2, nil)
	status, message, body, err := ctrl.Create(req)
	if err != nil {
		t.Fatalf(err.Error())
	}
	fmt.Printf("status:[%T]=> %v\n", status, status)
	fmt.Printf("message:[%T]=> %v\n", message, message)
	fmt.Printf("body:[%T]=> %v\n", body, body)
}

func TestCreate(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()
	sqlhandler := mock_database.NewMockSqlHandler(c)
	mockInteractor := mock_controllers.NewMockAccountInteractor(c)
	// ctx := context.Background()
	// kvs, err := store.NewKVS(ctx)
	// if err != nil {
	// 	log.Printf("[ERROR]: %+v", err)
	// }
	ctrl := controllers.NewAccountController(sqlhandler)
	ctrl.Interactor = mockInteractor

	cases := []struct {
		name             string
		requestBody      bool
		args             domain.Account
		prepareAddMockFn func(m *mock_controllers.MockAccountInteractor, args domain.Account)
		responseCode     int
		responseMessage  string
		responseBody     interface{}
		err              error
	}{
		{
			name:        "create = success",
			requestBody: true,
			args: domain.Account{
				Name:     "testUser",
				Email:    "test@exm.com",
				Password: "testPassword",
			},
			prepareAddMockFn: func(m *mock_controllers.MockAccountInteractor, args domain.Account) {
				m.EXPECT().Add(args).Return(
					&domain.Account{
						ID:       1,
						Name:     "testUser",
						Email:    "test@exm.com",
						Password: "testPassword",
					}, nil,
				)
			},
			responseCode:    200,
			responseMessage: "新規登録完了しました",
			responseBody: &domain.Account{
				ID:       1,
				Name:     "testUser",
				Email:    "test@exm.com",
				Password: "testPassword",
			},
			err: nil,
		},
		{
			name:        "when requestBody = nil, create = fail",
			requestBody: false,
			args:        domain.Account{},
			prepareAddMockFn: func(m *mock_controllers.MockAccountInteractor, args domain.Account) {
				m.EXPECT().Add(args).Return(nil, nil).AnyTimes()
			},
			responseCode:    400,
			responseMessage: "データ取得に失敗しました",
			responseBody:    nil,
			err:             errors.New("データ取得に失敗しました"),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			jsonArgs, _ := json.Marshal(tt.args)
			apiURL := "/api/register"
			if tt.requestBody {
				req = httptest.NewRequest("POST", apiURL, bytes.NewBuffer(jsonArgs))
			} else {
				req = httptest.NewRequest("POST", apiURL, nil)
			}
			tt.prepareAddMockFn(mockInteractor, tt.args)

			code, message, body, err := ctrl.Create(req)
			assert.Equal(t, tt.responseCode, code)
			assert.Equal(t, tt.responseMessage, message)
			assert.Equal(t, tt.responseBody, body)
			if err != nil {
				assert.Equal(t, tt.err.Error(), err.Error())
			}
		})
	}
}
