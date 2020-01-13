package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/gobuffalo/packr"
	"net/http"
)

func main() {

	box := packr.NewBox("../../assets/css")
	fmt.Println(box.FindString("application.css"))

	r := chi.NewRouter()
	distFileServer := http.StripPrefix("/css/", http.FileServer(box))
	r.Mount("/css/", distFileServer)
	http.ListenAndServe(":3000", r)





}

