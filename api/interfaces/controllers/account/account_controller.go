package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/kory-jp/vue_go/api/domain"
	"github.com/kory-jp/vue_go/api/interfaces/database"
	account "github.com/kory-jp/vue_go/api/interfaces/database/account"
	usecase "github.com/kory-jp/vue_go/api/usecase/account"
)

type AccountController struct {
	Interactor usecase.AccountInteractor
}

type Response struct {
	Status  int             `json:"status"`
	Message string          `json:"message"`
	Account *domain.Account `json:"account"`
}

func (res *Response) SetResp(status int, mess string, account *domain.Account) (resStr string) {
	response := &Response{status, mess, account}
	r, _ := json.Marshal(response)
	resStr = string(r)
	return
}

func NewAccountController(sqlHandler database.SqlHandler) *AccountController {
	return &AccountController{
		Interactor: usecase.AccountInteractor{
			AccountRepository: &account.AccountRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *AccountController) Create(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength == 0 {
		fmt.Println("NO DATA BODY")
		log.Println("NO DATA BODY")
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", nil)
		fmt.Fprintln(w, resStr)
		return
	}
	bytesAccount, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", nil)
		fmt.Fprintln(w, resStr)
		return
	}
	accountType := new(domain.Account)
	if err := json.Unmarshal(bytesAccount, accountType); err != nil {
		fmt.Println(err)
		log.Println(err)
		resStr := new(Response).SetResp(400, "データ取得に失敗しました", nil)
		fmt.Fprintln(w, resStr)
		return
	}
	account, err := controller.Interactor.Add(*accountType)
	if err != nil {
		errStr := err.Error()
		errStr1 := strings.Replace(errStr, "Error 1062: Duplicate entry", "入力された", 1)
		errStr2 := strings.Replace(errStr1, "for key 'email'", "既に登録されています。", 1)
		resStr := new(Response).SetResp(400, errStr2, nil)
		fmt.Println(err)
		log.Println(err)
		fmt.Fprintln(w, resStr)
		return
	}

	resStr := new(Response).SetResp(200, "新規登録完了しました", account)
	fmt.Fprintln(w, resStr)
}
