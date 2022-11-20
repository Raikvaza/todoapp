package main

import (
	"log"
	"net/http"

	"01.alem.school/git/aseitkha/ascii-art-web/server"
)

// func ErrorPage(w http.ResponseWriter, s string) error {
// 	errorPage, err := template.ParseFiles("templates/error.tmpl")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return err
// 	}
// 	errExec := errorPage.Execute(w, s)
// 	if errExec != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return errExec
// 	}
// 	return nil
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "POST" && r.Method != "GET" {
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 		return
// 	}
// 	switch r.URL.Path {
// 	case "/":
// 		mainpage(w, r)
// 	default:
// 		notfound(w, r)
// 	}
// }

func main() {
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))
	http.HandleFunc("/", server.HomeHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
