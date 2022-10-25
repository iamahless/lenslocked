package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err := fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
	if err != nil {
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

//func pathHandler(w http.ResponseWriter, r *http.Request) {
//	switch r.URL.Path {
//	case "/":
//		homeHandler(w, r)
//	case "/contact":
//		contactHandler(w, r)
//	case "/faq":
//		faqHandler(w, r)
//	default:
//		http.Error(w, "Page not found!", http.StatusNotFound)
//	}
//}

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, "Page not found!", http.StatusNotFound)
	}
}

func main() {
	var router Router
	fmt.Println("Starting the server on :3000...")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		return
	}
}
