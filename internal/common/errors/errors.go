package common_errors

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/astrokiran/nimbus/internal/common/log"
	"github.com/astrokiran/nimbus/internal/common/response"
	"go.uber.org/zap"
)

var logger = log.GetLogger()

func Test() {

}

func ReportServerError(r *http.Request, err error) {
	var (
		message = err.Error()
		method  = r.Method
		url     = r.URL.String()
		trace   = string(debug.Stack())
	)

	requestAttrs := zap.Strings("request", []string{"method", method, "url", url})
	logger.Error(message, requestAttrs, zap.String("trace", trace))
}

func ErrorMessage(w http.ResponseWriter, r *http.Request, status int, message string, headers http.Header) {
	message = strings.ToUpper(message[:1]) + message[1:]

	err := response.JSONWithHeaders(w, status, map[string]string{"Error": message}, headers)
	if err != nil {
		ReportServerError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func ServerError(w http.ResponseWriter, r *http.Request, err error) {
	ReportServerError(r, err)

	message := "The server encountered a problem and could not process your request"
	ErrorMessage(w, r, http.StatusInternalServerError, message, nil)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	message := "The requested resource could not be found"
	ErrorMessage(w, r, http.StatusNotFound, message, nil)
}

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("The %s method is not supported for this resource", r.Method)
	ErrorMessage(w, r, http.StatusMethodNotAllowed, message, nil)
}

// func (app *application) badRequest(w http.ResponseWriter, r *http.Request, err error) {
// 	errorMessage(w, r, http.StatusBadRequest, err.Error(), nil)
// }

// func (app *application) failedValidation(w http.ResponseWriter, r *http.Request, v validator.Validator) {
// 	err := response.JSON(w, http.StatusUnprocessableEntity, v)
// 	if err != nil {
// 		serverError(w, r, err)
// 	}
// }
