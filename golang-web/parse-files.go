package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type MFiles map[string]interface{}

func main() {
	http.HandleFunc("/parse-files", func(w http.ResponseWriter, r *http.Request) {
		var data = MFiles{"name": "Batman"}
		var tmpl = template.Must(template.ParseFiles(
			"views/parse-files.html",
			"views/_header.html",
			"views/_message.html",
		))

		var err = tmpl.ExecuteTemplate(w, "parse-files", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		var data = MFiles{"name": "Batman"}
		var tmpl = template.Must(template.ParseFiles(
			"views/about.html",
			"views/_header.html",
			"views/_message.html",
		))

		var err = tmpl.ExecuteTemplate(w, "about", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
