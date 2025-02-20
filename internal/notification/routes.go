package notification

import "github.com/go-chi/chi/v5"

func (n *Notification) NotificationRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/stream", n.StreamNotifications)

	return r

}
