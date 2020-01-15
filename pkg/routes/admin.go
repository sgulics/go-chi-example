package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/sgulics/go-chi-example/pkg/templating"
	"net/http"
)

func AdminRoutes(tm *templating.TemplateManager) chi.Router {
	//tmpl := template.Must(template.ParseFiles("templates_oild/index.gohtml"))
	//box := rice.MustFindBox("../../assets/css")

	//box := packr.NewBox("../../assets/css")

	//fmt.Println(box.FindString("application.css"))

	r := chi.NewRouter()

	//distFileServer := http.StripPrefix("/admin/css/", http.FileServer(box))
	//r.Mount("/css/", distFileServer)

	//basePath := "/assets"

	//r.Route(basePath, func(root chi.Router) {
	//	workDir, _ := os.Getwd()
	//	filesDir := filepath.Join(workDir, "public", "assets")
	//	FileServer(root, basePath, "/", http.Dir(filesDir))
	//})



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

//func FileServer(r chi.Router, basePath string, path string, root http.FileSystem) {
//	if strings.ContainsAny(path, "{}*") {
//		panic("FileServer does not permit URL parameters.")
//	}
//
//	fs := http.StripPrefix(basePath+path, http.FileServer(root))
//
//	if path != "/" && path[len(path)-1] != '/' {
//		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
//		path += "/"
//	}
//	path += "*"
//
//	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		fs.ServeHTTP(w, r)
//	}))
//}

