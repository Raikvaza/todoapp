package server

import (
	"log"
	"net/http"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(404)
		temp, err := template.ParseFiles("templates/error.tmpl")
		if err != nil {
			log.Fatal(err)
			return
		}
		err = temp.Execute(w, "404 Page Not Found")
		if err != nil {
			log.Fatal(err)
			return
		}
		return
	}
	if r.Method != http.MethodGet {
		w.WriteHeader(405)
		temp, err := template.ParseFiles("templates/error.tmpl")
		if err != nil {
			log.Fatal(err)
			return
		}
		err = temp.Execute(w, "405 Method Not Allowed")
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	html, err := template.ParseFiles("templates/base.tmpl")
	if err != nil {
		w.WriteHeader(500)
		temp, err := template.ParseFiles("templates/error.tmpl")
		if err != nil {
			log.Fatal(err)
			return
		}
		err = temp.Execute(w, "500 Internal Server Error")
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	err = html.Execute(w, html)
	if err != nil {
		w.WriteHeader(500)
		temp, err := template.ParseFiles("templates/error.tmpl")
		if err != nil {
			log.Fatal(err)
			return
		}
		err = temp.Execute(w, "500 Internal Server Error")
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}
