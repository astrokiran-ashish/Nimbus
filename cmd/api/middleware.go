package main

import (
	"fmt"
	"net/http"

	"github.com/astrokiranashish/nimbus/internal/common/response"
	"go.uber.org/zap"

	"github.com/tomasen/realip"
)

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				app.serverError(w, r, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func (app *application) logAccess(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mw := response.NewMetricsResponseWriter(w)
		next.ServeHTTP(mw, r)

		var (
			ip     = realip.FromRequest(r)
			method = r.Method
			url    = r.URL.String()
			proto  = r.Proto
		)

		userAttrs := zap.Strings("users", []string{"ip", ip})
		requestAttrs := zap.Strings("request", []string{"method", method, "url", url, "proto", proto})
		responseAttrs := zap.Strings("repsonse", []string{"status", fmt.Sprint(mw.StatusCode), "size", fmt.Sprint(mw.BytesCount)})

		app.logger.Info("access", userAttrs, requestAttrs, responseAttrs)
	})
}
