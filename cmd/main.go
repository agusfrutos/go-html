package main

import (
	"html/template"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("view/home.html")

	//w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, nil)
}

func main() {

	http.HandleFunc("/", homeHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Corriendo en http://localhost:3000/")

	log.Fatal(http.ListenAndServe(":3000", nil))
}
