package consultant

import "github.com/go-chi/chi/v5"

func (c *Consultant) ConsultantRoutes() chi.Router {
	r := chi.NewRouter()

	r.Post("/login", c.Login)
	r.Get("/", c.GetConsultant)

	return r

}
