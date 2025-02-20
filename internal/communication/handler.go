package communication

import (
	"encoding/json"
	"net/http"

	common_errors "github.com/astrokiran/nimbus/internal/common/errors"
	"github.com/astrokiran/nimbus/internal/models/nimbus/public/model"
)

func (com *Communication) WebhookEvent(w http.ResponseWriter, r *http.Request) {
	var event *model.AgoraWebhookEvents
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		common_errors.ErrorMessage(w, r, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

}
