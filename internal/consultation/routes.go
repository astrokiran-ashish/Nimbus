package consultation

import "github.com/go-chi/chi/v5"

func (con *Consultation) ConsultationRoutes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", con.CreateConsultationHandler)
	r.Get("/{consultationID}", con.GetConsultatioHandler)
	r.Put("/{consultationID}", con.UpdateConsultationHandler)
	r.Post("/{consultationID}/action", con.ConsultantActionEventHandler)
	r.Post("/events", con.ConsultationSessionEvents)

	return r

}
