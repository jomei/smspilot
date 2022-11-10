package sms

import (
	"SmsPilot2/pkg/jwt"
	"SmsPilot2/repository/smsDB/queueDB"
	"SmsPilot2/repository/smsDB/templateDB"
	"SmsPilot2/repository/usersDb"
	"encoding/json"
	"fmt"
	"github.com/rivo/uniseg"
	"net/http"
	"time"
)

type Sms struct {
	Phone              string            `json:"phone"`
	Template           int               `json:"template"`
	Macros             map[string]string `json:"macros"`
	DataTimeSendUTC    string            `json:"dataTimeSendUTC"`
	DataTimeFormatTime time.Time
}

type YourResponse struct {
	Id              int    `json:"id,omitempty"`
	Message         string `json:"message,omitempty"`
	Phone           int    `json:"phone,omitempty"`
	DataTimeSending string `json:"dataTimeSending,omitempty"`
}

func Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	// почитаем из заголовка токен и расшифруем его
	token := r.Header.Get("Token")
	decodeToken, err := jwt.Decode(token)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(decodeToken.IdUser)

	// Получаем на входе JSON добавления sms и конвертируем его в структуру.
	var smsObj Sms
	err = json.NewDecoder(r.Body).Decode(&smsObj)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "your format JSON is wrong")
		return
	}

	fmt.Println(smsObj)

	// Проверить
	// 1. в базе данных: активный данный юзер или нет (TODO: добавить кеширование на X минут)
	if !usersDb.IsUserActive(decodeToken.IdUser) {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "user is not active")
		return
	}

	// 2. есть ли такой шаблон в базе данных, если есть вернуть текст шаблона
	template, err := templateDB.GetTemplateById(smsObj.Template)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "id template is bad")
		return
	}

	// 3. есть все значения макросов если есть макросы
	err = CheckMarcos(smsObj, template)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "macros "+err.Error()+" not found in your request")
		return
	}

	// 4. указано правильно дата и время.
	fmt.Println(smsObj.DataTimeSendUTC)
	t, err := time.Parse("2006-01-02 15:04", smsObj.DataTimeSendUTC)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, smsObj.DataTimeSendUTC+" - time and date are not correct, format: 2006-01-02 15:04")
		return
	}
	smsObj.DataTimeFormatTime = t

	// 5. проверка верно ли указан номер телефона, вернем int номер телефона если верно
	phoneInt, err := PhoneGet(smsObj.Phone)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "bad format phone number")
		return
	}

	// Генерация смс текста
	jobToAdd := GeneratorMessage(smsObj, template, phoneInt, decodeToken.IdUser)
	fmt.Println(jobToAdd)

	// проверим длинну полной смс
	if uniseg.GraphemeClusterCount(jobToAdd.Message) > 160 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "your message is more than 160 characters")
		return
	}

	// добавление в очередь
	result := queueDB.AddJobToQueue(jobToAdd)
	var yourResponse YourResponse
	yourResponse.Id = result.Id
	yourResponse.Phone = result.Phone
	yourResponse.DataTimeSending = result.DatetimeSend

	yourResponseJson, err := json.Marshal(yourResponse)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, string(yourResponseJson))

}
