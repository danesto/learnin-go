package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type PageData struct {
	PageTitle string
}

func renderTemplate(w http.ResponseWriter, tmpl string, data PageData) {
	templates, err := template.ParseFiles("layout/layout.html", "layout/header.html", "layout/footer.html", tmpl)
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	err = templates.ExecuteTemplate(w, "layout.html", data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func main() {
	port := ":8080"

	fmt.Println("Hello, World!")

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{PageTitle: "About Page"}
		renderTemplate(w, "about.html", data)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "index.html", PageData{PageTitle: "Home page"})
	})

	fmt.Printf("server started at %s", port)
	http.ListenAndServe(port, nil)
}
