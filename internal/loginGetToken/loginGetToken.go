package loginGetToken

import (
	"SmsPilot2/pkg/jwt"
	"SmsPilot2/pkg/reqJSONToStruct"
	"SmsPilot2/repository/usersDb"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type TokenData struct {
	Token          string
	ExpirationDate int64
}

func Get(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// декодируем запрос юзера JSON в структуру
	var userDB usersDb.UserDB
	err := reqJSONToStruct.Convert(request, &userDB)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, "username or password is incorrect")
		return
	}

	// проверить есть в базе данных
	if !usersDb.IsUserValid(&userDB) {
		w.WriteHeader(http.StatusForbidden)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, "username or password is invalid")
		return
	}
	// сделать токен
	expDate := time.Now().UTC().Add(time.Hour * 2160).Unix()

	token, err := jwt.Get(userDB.Id, expDate)
	if err != nil {
		return
	}
	tokenData := TokenData{
		Token:          token,
		ExpirationDate: expDate,
	}
	// нарисовать его JSON
	b, err := json.Marshal(tokenData)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, string(b))
}
