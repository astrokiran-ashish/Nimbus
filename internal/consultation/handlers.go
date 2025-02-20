package consultation

import (
	"encoding/json"
	"fmt"
	"net/http"

	common_errors "github.com/astrokiran/nimbus/internal/common/errors"
	"github.com/astrokiran/nimbus/internal/models/nimbus/public/model"
	"github.com/google/uuid"
)

func (con *Consultation) CreateConsultationHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateConsultationRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		common_errors.ErrorMessage(w, r, http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	var consultation *model.Consultation
	if consultation, err = con.HandleConsultationCreation(req); err != nil {
		common_errors.ErrorMessage(w, r, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	if err := json.NewEncoder(w).Encode(consultation); err != nil {
		common_errors.ErrorMessage(w, r, http.StatusInternalServerError, "Failed to encode response", nil)
		return
	}
}

func (con *Consultation) GetConsultatioHandler(w http.ResponseWriter, r *http.Request) {
	consultationIDStr := r.URL.Query().Get("consultation_id")
	consultationID, err := uuid.Parse(consultationIDStr)
	if err != nil {
		common_errors.ErrorMessage(w, r, http.StatusBadRequest, "Invalid consultation ID", nil)
		return
	}

	var consultation *model.Consultation
	if consultation, err = con.HandleGetConsultation(consultationID); err != nil {
		common_errors.ErrorMessage(w, r, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	if err := json.NewEncoder(w).Encode(consultation); err != nil {
		common_errors.ErrorMessage(w, r, http.StatusInternalServerError, "Failed to encode response", nil)
		return
	}
}

func (con *Consultation) UpdateConsultationHandler(w http.ResponseWriter, r *http.Request) {
	var req UpdateConsultationRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		common_errors.ErrorMessage(w, r, http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	var consultation *model.Consultation
	if consultation, err = con.HandleUpdateConsultation(req); err != nil {
		common_errors.ErrorMessage(w, r, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	if err := json.NewEncoder(w).Encode(consultation); err != nil {
		common_errors.ErrorMessage(w, r, http.StatusInternalServerError, "Failed to encode response", nil)
		return
	}
}

func (con *Consultation) ConsultantActionEventHandler(w http.ResponseWriter, r *http.Request) {
	var req ConsultantActionEventRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		common_errors.ErrorMessage(w, r, http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	if err := con.HandleConsultantActionEvent(req); err != nil {
		common_errors.ErrorMessage(w, r, http.StatusInternalServerError, err.Error(), nil)
		return
	}

}

func (con *Consultation) ConsultationSessionEvents(w http.ResponseWriter, r *http.Request) {
	var event WebhookRequest
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		common_errors.ErrorMessage(w, r, http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	fmt.Printf("eventType %d eventChannelName %s ", event.EventType, event.Payload.ChannelName)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Ok"))
}
