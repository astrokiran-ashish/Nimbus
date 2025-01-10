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

type NimbusHTTPErrors struct {
}

var logger = log.GetLogger()

func (app *NimbusHTTPErrors) ReportServerError(r *http.Request, err error) {
	var (
		message = err.Error()
		method  = r.Method
		url     = r.URL.String()
		trace   = string(debug.Stack())
	)

	requestAttrs := zap.Strings("request", []string{"method", method, "url", url})
	logger.Error(message, requestAttrs, zap.String("trace", trace))
}

func (app *NimbusHTTPErrors) ErrorMessage(w http.ResponseWriter, r *http.Request, status int, message string, headers http.Header) {
	message = strings.ToUpper(message[:1]) + message[1:]

	err := response.JSONWithHeaders(w, status, map[string]string{"Error": message}, headers)
	if err != nil {
		app.ReportServerError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *NimbusHTTPErrors) ServerError(w http.ResponseWriter, r *http.Request, err error) {
	app.ReportServerError(r, err)

	message := "The server encountered a problem and could not process your request"
	app.ErrorMessage(w, r, http.StatusInternalServerError, message, nil)
}

func (app *NimbusHTTPErrors) NotFound(w http.ResponseWriter, r *http.Request) {
	message := "The requested resource could not be found"
	app.ErrorMessage(w, r, http.StatusNotFound, message, nil)
}

func (app *NimbusHTTPErrors) MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("The %s method is not supported for this resource", r.Method)
	app.ErrorMessage(w, r, http.StatusMethodNotAllowed, message, nil)
}

// func (app *application) badRequest(w http.ResponseWriter, r *http.Request, err error) {
// 	app.errorMessage(w, r, http.StatusBadRequest, err.Error(), nil)
// }

// func (app *application) failedValidation(w http.ResponseWriter, r *http.Request, v validator.Validator) {
// 	err := response.JSON(w, http.StatusUnprocessableEntity, v)
// 	if err != nil {
// 		app.serverError(w, r, err)
// 	}
// }
