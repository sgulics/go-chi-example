package routes

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/gorilla/sessions"
	"github.com/sgulics/go-chi-example/pkg/templating"
	"html/template"
	"net/http"
)

var sessionName = "admin-session"
var store = sessions.NewCookieStore([]byte("something-very-secret"))

func adminFuncMap(w http.ResponseWriter, r *http.Request) template.FuncMap {
	session, _ := store.Get(r, sessionName)
	flashes := session.Flashes()
	err := session.Save(r, w)
	if err != nil {
		fmt.Println(err)
	}
	tm := template.FuncMap{
		"isLoggedIn": func() bool {
			loggedIn, _ := isLoggedIn(r)
			return loggedIn
		},
		"flashMessages": func() []interface{} {
			return flashes
		},
	}
	return tm
}

func isLoggedIn(r *http.Request) (bool, error){
	session, err := store.Get(r, sessionName)
	if err != nil {
		return false, err
	}
	id := session.Values["userId"]
	return id != nil, nil

}

func LoginCheck(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if r.RequestURI == "/admin/login" || r.RequestURI == "/admin/logout" {
			next.ServeHTTP(w, r.WithContext(ctx))
		}

		loggedIn, err := isLoggedIn(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}
		if !loggedIn {
			http.Redirect(w, r, "/admin/login", http.StatusFound)
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}


func AdminRoutes(tm *templating.TemplateManager) chi.Router {

	r := chi.NewRouter()
	r.Use(LoginCheck)

	renderPage := func (templateName string, w http.ResponseWriter, req *http.Request, data interface{}) {
		t, err := tm.Template(templateName)
		if err != nil {
			render.Render(w, req, ErrInvalidRequest(err))
			return
		}
		t.Funcs(adminFuncMap(w,req)).Execute(w, data)
	}

	r.Get("/login", func(w http.ResponseWriter, req *http.Request) {
		//_ = tm.RenderTemplate(w, "articles/index.gohtml", nil)
		session, err := store.Get(req, sessionName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		delete(session.Values, "userId")

		session.AddFlash("Successfully Logged Out")
		session.Save(req, w)
		renderPage("login", w, req, nil)
	})

	r.Post("/login", func(w http.ResponseWriter, req *http.Request) {

		path := req.URL.Path
		session, err := store.Get(req, sessionName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		session.Values["userId"] = "123"
		session.AddFlash("Welcome to the Admin app!")
		session.Save(req, w)
		http.Redirect(w, req, path[:len(path)-5], 301)
	})

	r.Get("/logout", func(w http.ResponseWriter, req *http.Request) {

		//path := req.URL.Path
		//w.Write([]byte("LOGOUT.."))
		http.Redirect(w, req, "/admin/login", 301)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		renderPage("index", w, r, nil)

	})
	r.Get("/accounts", func(w http.ResponseWriter, r *http.Request) {
		renderPage("accounts/index", w, r, nil)
	})
	r.Get("/users/{userId}", func(w http.ResponseWriter, r *http.Request) {
		type myStruct struct {
			ID string
		}
		s := &myStruct{ID: chi.URLParam(r, "userId")}
		renderPage("users/show", w, r, s)

	})
	return r
}


