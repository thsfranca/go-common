package httphelper

import (
	"encoding/json"
	"errors"
	"github.com/ch4rl1e5/core/logger"
	"go.uber.org/zap"
	"net/http"
)

var ErrInvalidUUID = errors.New("invalid uuid")

type HTTPError struct {
	Description string // error description
	Status      int    `example:"500"` // status code (e.g. 500,404,400,200,201)
}

var MapErrors = map[string]HTTPError{}

func HandleError(w http.ResponseWriter, err error) {
	httpError := registeredError(err)
	http.Error(w, httpError.Description, httpError.Status)
}

func JsonResponse(w http.ResponseWriter, response interface{}) {
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		logger.ZapLogger.Panic("error:", zap.Error(err))
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		logger.ZapLogger.Panic("error:", zap.Error(err))
		return
	}
}

func registeredError(err error) HTTPError {
	return MapErrors[err.Error()]
}

func RegisterErrors(httpErrors ...HTTPError) {
	MapErrors[ErrInvalidUUID.Error()] = HTTPError{Description: ErrInvalidUUID.Error(), Status: http.StatusBadRequest}
	for _, v := range httpErrors {
		MapErrors[v.Description] = v
	}
}