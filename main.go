package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tplPath := filepath.Join("templates", "home.gohtml")

	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		log.Printf("Parsing template: %v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("Executing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err := fmt.Fprint(w, ""+
		"<h1>Contact Page</h1>"+
		"<p>To get in touch, please email me at "+
		"<a href=\"mailto:alexandergaruba96@gmail.com\">alexandergaruba96@gmail.com</a>."+
		"</p>")
	if err != nil {
		return
	}
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err := fmt.Fprint(w, "<h1>FAQ Page</h1>"+
		"<ul>"+
		"<li><b>Is there a free version?</b>Yes! We offer a free trial for 30 days on any paid plans.</li>"+
		"<li><b>What are your support hours?</b>We have support staff answering emails 24/7, though response times may be a bit slower on weekends.</li>"+
		"<li><b>How do I contact support?</b>Email me - <a href=\"mailto:alexandergaruba96@gmail.com\">alexandergaruba96@gmail.com</a></li></ul>")
	if err != nil {
		return
	}
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found!", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		return
	}
}
