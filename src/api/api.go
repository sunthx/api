package api

import (
	"encoding/json"
	"net/http"
	"tools"
)

func Guid(writer http.ResponseWriter, request *http.Request) {
	newGUID := tools.NewUUID()

	value, err := json.Marshal(&newGUID)
	if err != nil {
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Write(value)
}
