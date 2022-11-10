package reqJSONToStruct

import (
	"SmsPilot2/pkg/loggme"
	"encoding/json"
	"net/http"
)

// декодируем запрос юзера JSON в структуру

func Convert(r *http.Request, obj interface{}) error {
	err := json.NewDecoder(r.Body).Decode(&obj)
	loggme.Log.PrintAndWriteError("Convert(r *http.Request) - ", 1, 0, err)
	return err
}
