package response

import (
	"encoding/json"
	"github.com/go-http-utils/headers"
	"net/http"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   *error `json:"error"`
}

func CreateHandlerResponse(writer http.ResponseWriter, code int, message string, err *error) {
	writer.Header().Add(headers.ContentType, "application/json")
	writer.WriteHeader(code)
	errJon := json.NewEncoder(writer).Encode(Error{
		Code:    code,
		Message: message,
		Error:   err,
	})
	if errJon != nil {
		return
	}
}
