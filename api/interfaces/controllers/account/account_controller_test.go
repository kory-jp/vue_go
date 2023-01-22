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
	defer mock.EXPECT().Close().Return(nil)
	sqlhandler.EXPECT().Query(mysql.GetNumberAccountState, ac.Email).Return(mock, err)
	// --- FindById ---
	defer mock2.EXPECT().Close().Return(nil)
	mock2.EXPECT().Next().Return(true)
	mock2.EXPECT().Scan(&id, &name, &email, &password, &created_at).Return(nil)
	sqlhandler.EXPECT().Query(mysql.FindAccountState, 1).Return(mock2, nil)
	status, message, body, err := ctrl.Create(req)
	if err != nil {
		t.Fatalf(err.Error())
	}
	fmt.Printf("status:[%T]=> %v\n", status, status)
	fmt.Printf("message:[%T]=> %v\n", message, message)
	fmt.Printf("body:[%T]=> %v\n", body, body)
}
