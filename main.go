package main

import (
	//"context"
	//"errors"
	"flag"
	"fmt"
	//"math/rand"
	"net/http"
	//"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/docgen"
	//"github.com/go-chi/render"
)

var routes = flag.Bool("routes", false, "Generate router documentation")

func main() {
	flag.Parse()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	if *routes {
		//fmt.Println(docgen.JSONRoutesDoc(r))
		fmt.Println(docgen.MarkdownRoutesDoc(r, docgen.MarkdownOpts{
			ProjectPath: "github.com/josephbales/plain-budget-api",
			Intro:       "Welcome to the plain-budget-api generated docs.",
		}))
		return
	}

	http.ListenAndServe(":3333", r)
}
