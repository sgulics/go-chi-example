package routes

import (
	"github.com/go-chi/chi"
	"net/http"
)

func MonitorRoutes() chi.Router {
	r := chi.NewRouter()
	//r.Use(AdminOnly)
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	return r
}
