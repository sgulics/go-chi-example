package routes

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/gobuffalo/packr"
	"github.com/sgulics/go-chi-example/pkg/templating"
	"net/http"
)

func AdminRoutes(tm *templating.TemplateManager) chi.Router {
	//tmpl := template.Must(template.ParseFiles("templates_oild/index.gohtml"))
	//box := rice.MustFindBox("../../assets/css")

	box := packr.NewBox("../../assets/css")

	fmt.Println(box.FindString("application.css"))

	r := chi.NewRouter()

	distFileServer := http.StripPrefix("/admin/css/", http.FileServer(box))
	r.Mount("/css/", distFileServer)



	//r.Use(AdminOnly)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		//_ = tm.RenderTemplate(w, "articles/index.gohtml", nil)
		err := tm.Render(w, "index", nil)
		if err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}
	})
	r.Get("/accounts", func(w http.ResponseWriter, r *http.Request) {
		err := tm.Render(w, "accounts/index", nil)
		if err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}
	})
	r.Get("/users/{userId}", func(w http.ResponseWriter, r *http.Request) {
		type myStruct struct {
			ID string
		}
		s := &myStruct{ID: chi.URLParam(r, "userId")}
		err := tm.Render(w, "users/show", s )
		if err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}
	})
	return r
}