package main

import "net/http"
import "fmt"
import "html/template"

type Superhero struct {
	Name string
	Alias string
	Friends []string
}

func (s Superhero) SayHello(from string, message string) string {
	return fmt.Sprintf("%s said: \"%s\"", from, message)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var person = Superhero{
			Name: "Bruce Wayne",
			Alias: "Batman",
			Friends: []string{"Superman", "Flash", "Green Latern"},
		}

		var tmpl = template.Must(template.ParseFiles("views/view.html"))
		if err := tmpl 
	})
}