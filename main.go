package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Session struct {
	Make      string
	Model     string
	Watthours string
}

func main() {
	fmt.Println("Go htmx sample starting")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		sessions := map[string][]Session{
			"Sessions": {
				{Make: "Ford", Model: "F-150", Watthours: "420"},
				{Make: "Telsa", Model: "Model S", Watthours: "385"},
				{Make: "Kia", Model: "EV6", Watthours: "405"},
			},
		}
		tmpl.Execute(w, sessions)
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		make := r.PostFormValue("make")
		model := r.PostFormValue("model")
		watthours := r.PostFormValue("watthours")
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "session-list-element", Session{Make: make, Model: model, Watthours: watthours})
	}

	// define handlers
	http.HandleFunc("/", h1)
	http.HandleFunc("/add-session/", h2)

	log.Fatal(http.ListenAndServe(":8000", nil))

}
