package routes

import (
	"github.com/go-chi/chi"
	"net/http"
)



func AssetRoutes() chi.Router {

	r := chi.NewRouter()
	FileServer(r)
	return r
}


func FileServer(router *chi.Mux) {
	root := "public/assets"
	fs := http.FileServer(http.Dir(root))
	distFileServer := http.StripPrefix("/assets/", fs)
	router.Mount("/", distFileServer)

}