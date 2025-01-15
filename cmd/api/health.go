package main

import (
	"net/http"

	commonErrors "github.com/astrokiran/nimbus/internal/common/errors"

	"github.com/astrokiran/nimbus/internal/common/response"
)

func (app *application) status(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Status": "OK",
	}

	err := response.JSON(w, http.StatusOK, data)
	if err != nil {
		commonErrors.ServerError(w, r, err)
	}
}
