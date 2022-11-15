package server

import (
	"fmt"
	"net/http"
	"text/template"

	"01.alem.school/git/aseitkha/ascii-art-web/raiko"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Page not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	html, err := template.ParseFiles("templates/base.tmpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = html.Execute(w, html)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii" {
		http.Error(w, "404 Page not found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		http.Error(w, "405 Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	switch r.Method {
	case http.MethodPost: // POST /ASCII
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		text := r.FormValue("message")
		fmt.Println("Typed text : ", text)

		button := r.FormValue("banner")
		fmt.Println("Checked banner : ", button)

		if button != "standard" && button != "shadow" && button != "thinkertoy" {
			http.Error(w, "400 bad request", http.StatusBadRequest)
			return
		}
		html, err := template.ParseFiles("templates/ascii.tmpl")
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		data, err := raiko.Run(text, button)
		if err != nil { // Error handling Ascii art
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		err = html.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodGet: // GET /ASCII
		html, err := template.ParseFiles("templates/ascii.tmpl")
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		data, err := raiko.Run("No data specified", "standard")
		if err != nil { // Error handling Ascii art
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		err = html.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
