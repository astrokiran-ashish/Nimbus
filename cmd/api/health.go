package main

import (
	"net/http"

	"github.com/astrokiran/nimbus/internal/common/response"
)

func (app *application) status(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Status": "OK",
	}

	err := response.JSON(w, http.StatusOK, data)
	if err != nil {
		app.commonErrors.ServerError(w, r, err)
	}
}
