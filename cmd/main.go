package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./view/home.html"))
		tmpl.Execute(w, nil)
	})
	log.Println("Corriendo en puerto: 3000")

	log.Fatal(http.ListenAndServe(":3000", nil))
}
